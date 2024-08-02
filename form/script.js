// Generate random letters and animate them
const lettersContainer = document.getElementById('background-letters');

const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
const numLetters = 50;

for (let i = 0; i < numLetters; i++) {
    const letter = document.createElement('div');
    letter.className = 'letter';
    letter.textContent = letters[Math.floor(Math.random() * letters.length)];

    // Randomize position and animation delay
    letter.style.left = `${Math.random() * 100}vw`;
    letter.style.top = `${Math.random() * 100}vh`;
    letter.style.animationDelay = `${Math.random() * 5}s`;

    lettersContainer.appendChild(letter);
}
