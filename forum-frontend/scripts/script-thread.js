const overlay = document.querySelector(".overlay")
const deleteBtns = document.querySelectorAll(".delete-btn")
const deleteThreadBtn = document.querySelector(".delete-thread-btn")
const editBtns = document.querySelectorAll(".edit-btn")
const editThreadBtn = document.querySelector(".edit-thread-btn")
const addBtn = document.querySelector(".add-message-btn")

function displayAddMessageForm() {
    overlay.classList.toggle("invisible")
    document.querySelector(".add-message-form-container").classList.toggle("invisible")
}

function displayDeleteMessageForm() {
    overlay.classList.toggle("invisible")
    document.querySelector(".delete-message-form-container").classList.toggle("invisible")

    document.querySelector(".message-submit-btn-no").addEventListener("click", () => closeModal())
}

function displayDeleteThreadForm() {
    overlay.classList.toggle("invisible")
    document.querySelector(".delete-thread-form-container").classList.toggle("invisible")

    document.querySelector(".thread-submit-btn-no").addEventListener("click", () => closeModal())
}

function displayUpdateMessageForm() {
    overlay.classList.toggle("invisible")
    document.querySelector(".update-message-form-container").classList.toggle("invisible")

    document.querySelector(".message-submit-btn-no").addEventListener("click", () => closeModal())
}

function displayUpdateThreadForm() {
    overlay.classList.toggle("invisible")
    document.querySelector(".update-thread-form-container").classList.toggle("invisible")

    document.querySelector(".thread-submit-btn-no").addEventListener("click", () => closeModal())
}

function closeModal() {
    if (!(document.querySelector(".add-message-form-container").classList.contains("invisible"))) { document.querySelector(".add-message-form-container").classList.toggle("invisible") }
    if (!(document.querySelector(".delete-message-form-container").classList.contains("invisible"))) { document.querySelector(".delete-message-form-container").classList.toggle("invisible") }
    if (!(document.querySelector(".delete-thread-form-container").classList.contains("invisible"))) { document.querySelector(".delete-thread-form-container").classList.toggle("invisible") }
    if (!(document.querySelector(".update-message-form-container").classList.contains("invisible"))) { document.querySelector(".update-message-form-container").classList.toggle("invisible") }
    if (!(document.querySelector(".update-thread-form-container").classList.contains("invisible"))) { document.querySelector(".update-thread-form-container").classList.toggle("invisible") }
    overlay.classList.toggle("invisible")
}

addBtn.addEventListener("click", () => displayAddMessageForm())

if (deleteBtns) {
    deleteBtns.forEach( btn => {
        btn.addEventListener("click", () => {
            const messageId = btn.id
            const threadId = document.querySelector(".thread-infos-container").id
            document.querySelector(".delete-message-form-tag").action = `/delete/message/${threadId}&${messageId}`
            displayDeleteMessageForm()
        })
    })
}

if (deleteThreadBtn) { deleteThreadBtn.addEventListener("click", () => displayDeleteThreadForm()) }

if (editBtns) {
    editBtns.forEach( btn => {
        btn.addEventListener("click", () => {
            const messageId = btn.id
            const threadId = document.querySelector(".thread-infos-container").id
            document.querySelector(".update-message-form-tag").action = `/update/message/${threadId}&${messageId}`
            displayUpdateMessageForm()
        })
    })
}

if (editThreadBtn) { editThreadBtn.addEventListener("click", () => displayUpdateThreadForm()) }

overlay.addEventListener("click", (e) => {
    if (e.target === e.currentTarget) closeModal()
})