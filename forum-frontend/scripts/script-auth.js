document.querySelector(".register-redirection-btn").addEventListener( "click", () => {
    displayRegisterMenu()
})

function displayRegisterMenu() {
    document.querySelector(".menu-title").innerHTML = 'Créer un compte'

    const authForm = document.querySelector(".auth-form")

    authForm.innerHTML = `
        <input class="login-form-input" id="email" name="email" type="text" placeholder="Adresse email">
        <input class="login-form-input" id="username" name="username" type="text" placeholder="Nom d'utilisateur">
        <input class="login-form-input" id="password" name="password" type="text" placeholder="Mot de passe">
    `

    const registerRedirCont = document.querySelector(".register-redirection-container")

    registerRedirCont.innerHTML = `
        <p>Déjà un compte ?</p>
        <p class="register-redirection-btn">Se connecter</p>
    `

}