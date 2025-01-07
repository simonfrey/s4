/**
 * Unicode to ASCII (encode data to Base64)
 * @param {string} data
 * @return {string}
 */
function utoa(data) {
    return btoa(unescape(encodeURIComponent(data)));
}

function fillEncryptOutputs(values) {
    const output = document.getElementById("encryptOutput")
    output.innerHTML = ""

    values.forEach((value) => {
        const textarea = document.createElement("textarea")
        textarea.readOnly = true
        textarea.innerHTML = value
        output.appendChild(textarea)
    })
}

function doEncrypt() {

    const thresholdEl = document.getElementById("encryptThreshold")
    const sharesEl = document.getElementById("encryptShares")

    // Validation
    if (sharesEl.valueAsNumber <= 1) {
        sharesEl.value = 2
    }
    if (thresholdEl.valueAsNumber <= 1) {
        thresholdEl.value = 2
    }

    // Ensure threshold isn't higher than shares
    thresholdEl.setAttribute("max", sharesEl.valueAsNumber)
    if (sharesEl.valueAsNumber < thresholdEl.valueAsNumber) {
        thresholdEl.value = sharesEl.value
    }

    // Snag values from the DOM
    const threshold = thresholdEl.valueAsNumber
    const shares = sharesEl.valueAsNumber
    const useAES = document.getElementById("encryptUseAES").checked
    const useBase24 = document.getElementById("encryptUseBase24").checked
    const input = document.getElementById("encryptInput").value

    // Handle no input
    if (input === "") {
        return fillEncryptOutputs(Array(shares).fill(""))
    }
    const encoded = new TextEncoder().encode('€')
    // Do it!
    const res = Distribute_fours(
        utoa(input),
        shares,
        threshold,
        useAES,
        useBase24,
    )

    // Update DOM with results
    if (typeof res === "string") {
        setError(res)
        fillEncryptOutputs(Array(shares).fill(""))
    } else {
        setError("")
        fillEncryptOutputs(res)
    }
}
