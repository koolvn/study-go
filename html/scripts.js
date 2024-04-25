// Get the header and main elements
const header = document.querySelector('.header');
const main = document.querySelector('main');

// Add an event listener to detect when the user scrolls past the header
document.addEventListener('scroll', () => {
    // Check if the top of the window is below the top of the header
    if (window.scrollY >= header.offsetTop) {
        // If it is, add a class to make the header sticky
        header.classList.add('sticky');
    } else {
        // Otherwise, remove the sticky class
        header.classList.remove('sticky');
    }
});