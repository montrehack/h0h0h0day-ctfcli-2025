import sqlite3
import hashlib
import secrets

def init_database():
    conn = sqlite3.connect('santa.db')
    cursor = conn.cursor()
    
    # Drop existing tables
    cursor.execute('DROP TABLE IF EXISTS plebe_users')
    cursor.execute('DROP TABLE IF EXISTS gifts')
    cursor.execute('DROP TABLE IF EXISTS gift_assignments')
    cursor.execute('DROP TABLE IF EXISTS santa_users')
    
    print("✓ Dropped old tables")
    
    # Create plebe_users table
    cursor.execute('''
        CREATE TABLE plebe_users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        )
    ''')
    
    # Create gifts table
    cursor.execute('''
        CREATE TABLE gifts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT,
            value INTEGER
        )
    ''')
    
    # Create gift_assignments table
    cursor.execute('''
        CREATE TABLE gift_assignments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            elf_id INTEGER,
            gift_id INTEGER,
            FOREIGN KEY (elf_id) REFERENCES plebe_users(id),
            FOREIGN KEY (gift_id) REFERENCES gifts(id)
        )
    ''')
    
    # Create santa_users table
    cursor.execute('''
        CREATE TABLE santa_users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    username TEXT UNIQUE NOT NULL,
	    password_hash TEXT NOT NULL
        )
    ''')
    
    print("✓ Tables created")
    
    # Insert test plebe_users
    test_elves = [
    	('john', hashlib.md5('elfpwd321'.encode()).hexdigest()),
    	('sarah', hashlib.md5('elfpwd321'.encode()).hexdigest()),
    	('mike', hashlib.md5('elfpwd321'.encode()).hexdigest()),
	]
    
    for username, password_hash in test_elves:
        cursor.execute(
            'INSERT INTO plebe_users (username, password) VALUES (?, ?)',
            (username, password_hash)
        )
    
    print(f"✓ Inserted plebe_users")
    
    # Insert gifts
    gifts = [
        ('Candy Cane', 'Sweet treat', 5),
        ('Fidget spinner', 'No. Just no.', 7),
        ('Socks', 'Warm woolen socks', 10),
        ('Coffee Mug', 'Nice ceramic mug', 12),
        ('Book', 'Award winning novel', 20),
        ('Candle', 'Cozy peppermint perfume', 30),
        ('Headphones', 'To listen to your favorite music', 50),
        ('NVIDIA RTX 5090', 'Top-tier graphics card', 4000)
    ]
    
    for name, desc, value in gifts:
        cursor.execute(
            'INSERT INTO gifts (name, description, value) VALUES (?, ?, ?)',
            (name, desc, value)
        )
    
    print(f"✓ Inserted {len(gifts)} gifts")
    
    # Assign boring gifts to plebe
    cursor.execute('INSERT INTO gift_assignments (elf_id, gift_id) VALUES (1, 1)')  # john -> candy
    cursor.execute('INSERT INTO gift_assignments (elf_id, gift_id) VALUES (2, 2)')  # sarah -> socks
    cursor.execute('INSERT INTO gift_assignments (elf_id, gift_id) VALUES (3, 3)')  # mike -> mug
    
    print("✓ Assigned gifts to plebe")
    
    # Insert admin user
    # Password from rockyou.txt: "h0h0h0"
    # MD5 hash of "h0h0h0"
    admin_password = "h0h0h0"
    admin_hash = hashlib.md5(admin_password.encode()).hexdigest()
    
    cursor.execute(
        'INSERT INTO santa_users (username, password_hash) VALUES (?, ?)',
        ('santa_claus', admin_hash)
    )
    
    print(f"✓ Created admin user")
    
    conn.commit()
    conn.close()
    
    print("\n" + "="*60)
    print("DATABASE INITIALIZED SUCCESSFULLY!")

if __name__ == '__main__':
    init_database()
