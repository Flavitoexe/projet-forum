const threads = document.querySelectorAll(".thread-btn")

for (let i = 0; i < threads.length; i++) {
    threads[i].addEventListener("click", () => document.location.href = `/thread/${threads[i].id}`)
}
