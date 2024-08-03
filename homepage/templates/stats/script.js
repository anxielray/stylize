// Generate random letters and animate them
const lettersContainer = document.getElementById('background-letters');

const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
const numLetters = 20;

for (let i = 0; i < numLetters; i++) {
    const letter = document.createElement('div');
    letter.className = 'letter';
    letter.textContent = letters[Math.floor(Math.random() * letters.length)];

    // Randomize position and animation delay
    letter.style.left = `${Math.random() * 100}vw`;
    letter.style.top = `${Math.random() * 100}vh`;
    letter.style.fontFamily = `'Lucida Handwriting', cursive`;
    letter.style.fontSize = `${Math.random() * 2 + 1}rem`; // Random font size for variety
    letter.style.animationDelay = `${Math.random() * 5}s`;

    lettersContainer.appendChild(letter);
}


document.addEventListener('DOMContentLoaded', function () {
    const footer = document.querySelector('footer');
    const footerCreators = document.querySelector('.footer-creators');

    function onScroll() {
        const footerPosition = footer.getBoundingClientRect().top;
        const viewportHeight = window.innerHeight;

        if (footerPosition <= viewportHeight) {
            footerCreators.classList.add('show');
        } else {
            footerCreators.classList.remove('show');
        }
    }

    window.addEventListener('scroll', onScroll);
    // Check if the footer is in view on initial load
    onScroll();
});