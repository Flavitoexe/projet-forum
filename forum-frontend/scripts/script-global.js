document.querySelector(".logo-title").addEventListener("click", () => document.location.href = "/")

const registerBtn = document.querySelector(".register-btn")
const loginBtn = document.querySelector(".login-btn")

if (registerBtn) { registerBtn.addEventListener("click", () => document.location.href = "/register") }
if (loginBtn) { loginBtn.addEventListener("click", () => document.location.href = "/login") }

// Affichage des catégories populaires
const popularCategories = document.querySelectorAll(".popular-categories-name")

for (let i = 0; i < popularCategories.length; i++) {
    popularCategories[i].addEventListener("click", () => document.location.href = `/category/${popularCategories[i].id}`)
}

// Affichage sections 
document.querySelector(".populars").addEventListener("click", () => document.location.href = "/populars")
document.querySelector(".news").addEventListener("click", () => document.location.href = "/news")
document.querySelector(".explorer").addEventListener("click", () => document.location.href = "/categories")

// // Formulaire de thread
// document.querySelector(".add-thread-btn").addEventListener("click", () => displayAddThreadForm())
// const overlay = document.querySelector(".overlay")

// function displayAddThreadForm() {
//     overlay.classList.toggle("invisible")

//     const threadForm = document.createElement("div")
//     threadForm.className = "add-thread-form-container"

//     threadForm.innerHTML = `
//         <h2 class="add-thread-title">De quoi voulez vous discuter ?</h2>
//         <div class="add-thread-form">
//             <form action="/thread" method="POST">
//                 <input class="thread-form-input" id="content" name="content" type="text" placeholder="Blabla bla..." required>
//             </form>
//             <button type="submit" class="thread-submit-btn submit-btn">Envoyer</button>
//         </div>
//     `
//     document.querySelector(".overlay").appendChild(threadForm)

//     overlay.addEventListener("click", (e) => {
//         if (e.target === e.currentTarget) {
//             clearAddThreadForm()
//         }
//     })
// }

// function clearAddThreadForm() {
//     document.querySelector(".add-thread-form-container").remove()
//     overlay.classList.toggle("invisible")
// }