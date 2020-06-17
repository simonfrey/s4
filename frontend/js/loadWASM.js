window.wasmLoaded = false;

if (!WebAssembly.instantiateStreaming) {
    // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
let mod, inst;

var wasmResponse =
WebAssembly.instantiateStreaming(fetch(wasmPayloadInlineURL), go.importObject).then(
    async  result => {
        console.log("Loaded")
        mod = result.module;
        inst = result.instance;

        window.wasmLoaded = true;
        await go.run(inst);
        inst = await WebAssembly.instantiate(mod, go.importObject);
    }
).catch(err =>{
    console.log("Webassembly ERROR: ",err)
    errorNotification.error = err
});