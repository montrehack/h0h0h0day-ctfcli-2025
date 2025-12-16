document.addEventListener('DOMContentLoaded', () => {

    const messageArea = document.getElementById('message-area');
    const loginForm = document.getElementById('login-form');
    const registerForm = document.getElementById('register-form');

    function showMessage(type, message) {
        messageArea.innerHTML = `
            <div class="message-alert ${type}">
                ${message}
            </div>
        `;
        setTimeout(() => {
            messageArea.innerHTML = '';
        }, 5000);
    }

    async function handleAuthSubmission(event, endpoint) {
        event.preventDefault();

        const form = event.target;
        const formData = new FormData(form);
        const submitButton = form.querySelector('button[type="submit"]');

        const bodyData = new URLSearchParams(formData).toString(); 

        submitButton.disabled = true;
        submitButton.textContent = 'Processing...';

        try {
            const response = await fetch(endpoint, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                body: bodyData 
            });

            const data = await response.json(); 

            if (response.ok) {
                showMessage('success', data.message || 'Success! Redirecting...');
                
                setTimeout(() => {
                    window.location.replace(data.redirect || '/wishlist');
                }, 1000); 

            } else {
                showMessage('error', data.message || 'An unexpected error occurred.');
            }

        } catch (error) {
            console.error('Network Error:', error);
            showMessage('error', 'A network error occurred. Please check your connection.');
        } finally {
            submitButton.disabled = false;
            submitButton.textContent = (endpoint === '/login' ? 'Login' : 'Create Account');
        }
    }
    
    if (loginForm) {
        loginForm.addEventListener('submit', (e) => handleAuthSubmission(e, '/login'));
    }

    if (registerForm) {
        registerForm.addEventListener('submit', (e) => handleAuthSubmission(e, '/register'));
    }
});