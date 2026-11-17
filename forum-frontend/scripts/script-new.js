const newThreads = document.querySelectorAll(".thread-btn")

for (let i = 0; i < newThreads.length; i++) {
    newThreads[i].addEventListener("click", () => document.location.href = `/thread/${newThreads[i].id}`)
}