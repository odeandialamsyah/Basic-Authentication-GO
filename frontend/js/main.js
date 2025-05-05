// API Base URL
const API_BASE_URL = 'http://localhost:8080';

// DOM Elements
const loginTab = document.getElementById('loginTab');
const registerTab = document.getElementById('registerTab');
const loginForm = document.getElementById('loginForm');
const registerForm = document.getElementById('registerForm');
const alert = document.getElementById('alert');
const alertIcon = document.getElementById('alertIcon');
const alertMessage = document.getElementById('alertMessage');

// Tab Switching
loginTab.addEventListener('click', () => {
    loginForm.classList.remove('hidden');
    registerForm.classList.add('hidden');
    loginTab.classList.add('border-blue-500', 'text-blue-500');
    loginTab.classList.remove('border-gray-200', 'text-gray-500');
    registerTab.classList.add('border-gray-200', 'text-gray-500');
    registerTab.classList.remove('border-blue-500', 'text-blue-500');
});

registerTab.addEventListener('click', () => {
    registerForm.classList.remove('hidden');
    loginForm.classList.add('hidden');
    registerTab.classList.add('border-blue-500', 'text-blue-500');
    registerTab.classList.remove('border-gray-200', 'text-gray-500');
    loginTab.classList.add('border-gray-200', 'text-gray-500');
    loginTab.classList.remove('border-blue-500', 'text-blue-500');
});

// Show Alert Function
function showAlert(message, isSuccess = true) {
    alertMessage.textContent = message;
    alertIcon.className = isSuccess ? 'fas fa-check-circle text-green-500' : 'fas fa-exclamation-circle text-red-500';
    alert.classList.remove('hidden');
    setTimeout(() => {
        alert.classList.add('hidden');
    }, 3000);
}

// Login Form Handler
loginForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;

    try {
        const response = await fetch(`${API_BASE_URL}/auth/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });

        const data = await response.json();

        if (response.ok) {
            showAlert('Login successful!');
            // Store the token
            localStorage.setItem('token', data.token);
            // Redirect to dashboard or home page after successful login
            setTimeout(() => {
                window.location.href = '/dashboard.html';
            }, 1000);
        } else {
            showAlert(data.message || 'Login failed', false);
        }
    } catch (error) {
        showAlert('An error occurred. Please try again.', false);
    }
});

// Register Form Handler
registerForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const name = document.getElementById('registerName').value;
    const email = document.getElementById('registerEmail').value;
    const password = document.getElementById('registerPassword').value;

    try {
        const response = await fetch(`${API_BASE_URL}/auth/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username: name, email, password }),
        });

        const data = await response.json();

        if (response.ok) {
            showAlert('Registration successful! Please login.');
            // Switch to login form
            loginTab.click();
            registerForm.reset();
        } else {
            showAlert(data.message || 'Registration failed', false);
        }
    } catch (error) {
        showAlert('An error occurred. Please try again.', false);
    }
});
