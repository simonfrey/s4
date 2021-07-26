if (!WebAssembly.instantiateStreaming) {
    // Polyfill (e.g. Safari)
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer()
        return await WebAssembly.instantiate(source, importObject)
    }
}

function handleInternetConnectivity() {
    document.getElementById("notOffline").style.display = navigator.onLine ? "block" : "none"
}

function handleNotFileSaved() {
    document.getElementById("notFileSaved").style.display = window.location.protocol !== "file:" ? "block" : "none"
}

function setError(msg) {
    document.getElementById("error").style.display = msg ? "block" : "none";
    document.querySelector("#error p").innerHTML = msg;
}

// Clear textareas when pressing Ctrl+S to prevent saving secrets
// Does not work if the user save from the browser menu...
document.addEventListener("keydown", (e) => {
    if (e.ctrlKey && e.key === "s") {
        document.querySelectorAll("textarea").forEach((e) => e.value = "")
    }
})

document.addEventListener("DOMContentLoaded", async () => {
    // Dismissable alert
    document.querySelector("#error button").addEventListener("click", () => setError())
    // Tabs
    document.querySelectorAll("nav ul li").forEach((li) => {
        li.addEventListener("click", () => {
            document.querySelectorAll("nav ul li").forEach((el) => {
                const action = (li === el) ? "add" : "remove"
                el.classList[action]("active")
            })
            document.querySelectorAll("section").forEach((el) => {
                el.style.display = (el.id === li.dataset.section) ? "block" : "none"
            })
            setError("")
        })
    })
    // Security alerts
    window.addEventListener("online", handleInternetConnectivity)
    window.addEventListener("offline", handleInternetConnectivity)
    handleInternetConnectivity()
    handleNotFileSaved()
    // Load WASM file
    try {
        const go = new Go()
        const result = await WebAssembly.instantiateStreaming(fetch(wasmPayloadInlineURL), go.importObject)
        go.run(result.instance)
        // Load Encrypt
        document.querySelectorAll("#encryptThreshold, #encryptShares, #encryptUseAES, #encryptInput").forEach((input) => {
            input.addEventListener("input", doEncrypt)
        })
        // Load Decrypt
        document.getElementById("decryptShares").addEventListener("input", handleDecryptShareChange)
        handleDecryptShareChange()
    } catch (err) {
        setError(err)
    }
});
