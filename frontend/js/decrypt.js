function handleDecryptShareChange() {

    const inputsEl = document.getElementById("decryptInputs")
    const sharesEl = document.getElementById("decryptShares")

    let wanted = sharesEl.valueAsNumber
    const existing = inputsEl.childNodes.length

    // Validate
    if (wanted < 2) {
        wanted = 2
        sharesEl.value = 2
    }

    // Adjust the number of textareas
    if (wanted > existing) {
        for (let i = 0; i < wanted - existing; i++) {
            const ta = document.createElement("textarea")
            ta.placeholder = "Enter a share here"
            ta.addEventListener("input", doDecrypt)
            inputsEl.append(ta)
        }
    } else if (wanted < existing) {
        for (let i = 0; i < existing - wanted; i++) {
            inputsEl.removeChild(inputsEl.lastChild)
        }
    }
}

function doDecrypt() {

    // Collect inputs
    let inputs = []
    document.querySelectorAll("#decryptInputs > textarea").forEach((ta) => {
        inputs.push(ta.value)
    })

    // Check if they are all empty for no-op
    if (inputs.filter(Boolean).length === 0) {
        return
    }

    const res = Recover_fours(inputs)
    const outEl = document.getElementById("decryptOutput")
    const base64regex = /^([0-9a-zA-Z+/]{4})*(([0-9a-zA-Z+/]{2}==)|([0-9a-zA-Z+/]{3}=))?$/

    if (!base64regex.test(res)) {
        outEl.innerText = ""
        setError(res)
    } else {
        outEl.innerText = atob(res)
        setError("")
    }
}
