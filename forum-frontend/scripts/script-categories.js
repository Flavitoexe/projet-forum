const categories = document.querySelectorAll(".category-btn")

for (let i = 0; i < categories.length; i++) {
    categories[i].addEventListener("click", () => document.location.href = `/category/${categories[i].id}`)
}
