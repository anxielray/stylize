// JavaScript to toggle between showing the creator section with sliding effect and the username section
document.addEventListener('DOMContentLoaded', () => {
    const toggleButton = document.getElementById('toggleButton');
    const creatorSection = document.getElementById('creatorSection');
    const usernameSection = document.getElementById('usernameSection');

    toggleButton.addEventListener('click', () => {
        if (creatorSection.classList.contains('hidden')) {
            creatorSection.classList.remove('hidden');
            usernameSection.classList.add('hidden');
            toggleButton.textContent = 'Show Creators';
        } else {
            creatorSection.classList.add('hidden');
            usernameSection.classList.remove('hidden');
            toggleButton.textContent = 'Show Username';
        }
    });
});
