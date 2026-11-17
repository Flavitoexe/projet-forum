const threadBtn = document.querySelectorAll(".thread-btn")

for (let i = 0; i < threadBtn.length; i++) {
    threadBtn[i].addEventListener("click", () => document.location.href = `/thread/${threadBtn[i].id}`)
}
