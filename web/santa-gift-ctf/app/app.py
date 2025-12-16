from flask import Flask, render_template, request, redirect, url_for, session
import sqlite3
import hashlib

app = Flask(__name__)
# N'oubliez pas de mettre une vraie clé secrète !
app.secret_key = 'santa_secret_key_2024'

def get_db(readonly=False):
    """Initialise et retourne une connexion à la base de données."""
    if readonly:
        # Permet d'ouvrir le fichier en lecture seule (pas supporté sur tous les systèmes, mais bonne pratique)
        conn = sqlite3.connect('file:santa.db?mode=ro', uri=True)
    else:
        conn = sqlite3.connect('santa.db')
    conn.row_factory = sqlite3.Row
    return conn

@app.route('/')
def index():
    """Page d'accueil simple, redirige vers le tableau de bord si connecté."""
    if 'username' in session:
        return redirect(url_for('dashboard'))
    return redirect(url_for('login'))

@app.route('/login', methods=['GET', 'POST'])
def login():
    """Gère la connexion pour les utilisateurs Plebe et Santa."""
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        
        conn = get_db()
        cursor = conn.cursor()
        
        password_hash = hashlib.md5(password.encode()).hexdigest()
        user = cursor.execute('SELECT id, username FROM plebe_users WHERE username = ? AND password = ?', 
                              (username, password_hash)).fetchone()

        is_santa = False
        
        if not user:
            user = cursor.execute('SELECT id, username FROM santa_users WHERE username = ? AND password_hash = ?', 
                                  (username, password_hash)).fetchone()
            if user:
                is_santa = True

        conn.close()

        if user:
            session['username'] = user['username']
            session['user_id'] = user['id']
            session['is_santa'] = is_santa
            return redirect(url_for('dashboard'))
        else:
            return render_template('login.html', error="Invalid username or password.")

    return render_template('login.html')

@app.route('/register', methods=['GET', 'POST'])
def register():
    """Gère l'inscription des nouveaux utilisateurs Plebe et les connecte automatiquement."""
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']

        if len(username) < 3 or len(password) < 5:
            return render_template('register.html', error="Username must be at least 3 chars, password 5 chars.")

        conn = get_db()
        cursor = conn.cursor()
        
        try:
            # Hashage du mot de passe
            password_hash = hashlib.md5(password.encode()).hexdigest()

            # Insertion de l'utilisateur Plebe
            cursor.execute('INSERT INTO plebe_users (username, password) VALUES (?, ?)', 
                           (username, password_hash))
            new_user_id = cursor.lastrowid
            
            # Attribuer le cadeau par défaut (Candy Cane, ID 1)
            cursor.execute('INSERT INTO gift_assignments (elf_id, gift_id) VALUES (?, 1)', 
                           (new_user_id,))
            
            conn.commit()
            conn.close()
            
            # --- LOGIQUE D'AUTO-CONNEXION ---
            session['username'] = username
            session['user_id'] = new_user_id
            session['is_santa'] = False # Un utilisateur enregistré n'est jamais Santa
            
            # Redirige directement vers le tableau de bord
            return redirect(url_for('dashboard'))
        
        except sqlite3.IntegrityError:
            conn.close()
            return render_template('register.html', error="Username already exists.")

    return render_template('register.html')

@app.route('/dashboard')
def dashboard():
    """Tableau de bord affichant les assignations (Lecture seule)."""
    if 'username' not in session:
        return redirect(url_for('index'))
    
    username = session['username']
    is_santa = session.get('is_santa', False)
    
    conn = get_db(readonly=True)
    cursor = conn.cursor()

    if is_santa:
        # Santa view - show all assignments (READ-ONLY)
        assignments = cursor.execute('''
            SELECT e.username, e.id as elf_id, g.name as gift_name, g.id as gift_id
            FROM gift_assignments ga
            JOIN plebe_users e ON ga.elf_id = e.id
            JOIN gifts g ON ga.gift_id = g.id
        ''').fetchall()
        
        # ADD THIS: Get all gifts for Santa dashboard
        gifts = cursor.execute('SELECT * FROM gifts').fetchall()
        
        context = {
            'username': username,
            'is_santa': True,
            'assignments': assignments,
            'gifts': gifts,  # ADD THIS LINE
            'view': 'santa_read'
        }
    else:
        # Regular elf view - show only their gift
        user_id = session['user_id']
        
        gift = cursor.execute('''
            SELECT g.id, g.name, g.description, g.value
            FROM gift_assignments ga
            JOIN gifts g ON ga.gift_id = g.id
            WHERE ga.elf_id = ?
        ''', (user_id,)).fetchone()
        
        context = {
            'username': username,
            'is_santa': False,
            'gift': gift,
            'view': 'elf'
        }

    conn.close()
    return render_template('dashboard.html', **context)
@app.route('/assign_gift', methods=['GET', 'POST'])
def assign_gift():
    """Page d'administration pour l'assignation des cadeaux (Santa uniquement)."""
    if 'username' not in session:
        return redirect(url_for('index'))

    is_santa = session.get('is_santa', False)
    
    # L'Elfe est connecté mais n'est pas Santa. Affiche le message d'accès refusé.
    if not is_santa:
        return render_template('assign_gift.html', is_santa=False)

    # Logique pour Santa
    conn = get_db()
    cursor = conn.cursor()
    message = None

    if request.method == 'POST':
        # NOTE : Les entrées ne sont pas filtrées (Vulnérabilité SQLi T2)
        target_elf_id = request.form.get('elf_id')
        new_gift_id = request.form.get('gift_id')
        
        if target_elf_id and new_gift_id:
            try:
                # Requête de mise à jour Vulnérable : injection possible via target_elf_id ou new_gift_id
                cursor.execute(f"UPDATE gift_assignments SET gift_id = {new_gift_id} WHERE elf_id = {target_elf_id}")
                conn.commit()
                message = {'type': 'success', 'text': f"Gift ID {new_gift_id} successfully assigned to User ID {target_elf_id}."}
            except Exception as e:
                 message = {'type': 'error', 'text': f"Database error during assignment: {str(e)}"}
        else:
            message = {'type': 'error', 'text': "Missing Elf ID or Gift ID."}


    # Chargement des données pour l'affichage (GET ou POST après traitement)
    assignments = cursor.execute('''
        SELECT e.username, e.id as elf_id, g.name as gift_name, g.id as gift_id
        FROM gift_assignments ga
        JOIN plebe_users e ON ga.elf_id = e.id
        JOIN gifts g ON ga.gift_id = g.id
    ''').fetchall()
    
    gifts = cursor.execute('SELECT * FROM gifts').fetchall()
    
    conn.close()
    
    return render_template('assign_gift.html', 
                         is_santa=True,
                         assignments=assignments,
                         gifts=gifts,
                         message=message)

@app.route('/search', methods=['GET'])
def search():
    """Page de recherche de cadeaux (POINT D'ENTRÉE SQLi T1)."""
    if 'username' not in session:
        return redirect(url_for('index'))
    
    conn = get_db(readonly=True)
    cursor = conn.cursor()
    
    keyword = request.args.get('keyword', '')
    results = []
    error = None
    
    # Requête vulnérable à la SQL Injection (READ-ONLY)
    # If no keyword, show all gifts
    if keyword:
        sql_query = f"SELECT name, description, value FROM gifts WHERE name LIKE '%{keyword}%' OR description LIKE '%{keyword}%'"
    else:
        sql_query = "SELECT name, description, value FROM gifts"
    
    try:
        results = cursor.execute(sql_query).fetchall()
    except Exception as e:
        # Affiche l'erreur pour aider au débug de l'injection
        error = f"Database Error: {str(e)}"
    
    conn.close()
    
    return render_template('search.html', keyword=keyword, results=results, error=error)

@app.route('/logout')
def logout():
    """Déconnexion de l'utilisateur."""
    session.clear()
    return redirect(url_for('index'))

if __name__ == '__main__':
    # Mode debug=True pour vous aider à voir les Internal Server Errors (500)
    # Si le conteneur s'arrête immédiatement, essayez debug=False.
    app.run(host='0.0.0.0', port=5000, debug=True)
