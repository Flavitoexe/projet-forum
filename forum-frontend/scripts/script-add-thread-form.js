// Formulaire de thread
document.querySelector(".add-thread-btn").addEventListener("click", () => overlay.classList.toggle("invisible"))

const overlay = document.querySelector(".overlay")
overlay.addEventListener("click", (e) => {
    if (e.target === e.currentTarget) {
        overlay.classList.toggle("invisible")
    }
})