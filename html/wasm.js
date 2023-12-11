const WASM_URL = 'wasm.wasm';

var wasm;
let go;
let wasmModule = null;

// WebAssembly.Moduleを格納する処理
function loadWasmModule() {
    fetch(WASM_URL).then(response =>
        response.arrayBuffer()
    ).then(bytes =>
        wasmModule = new WebAssembly.Module(bytes)
    );
};

function model_inference(input, rows, cols) {
    go = new Go();
    wasm = (new WebAssembly.Instance(wasmModule, go.importObject));
    go.run(wasm);

    const startTime = performance.now();
    let pred = window.inference(input, rows, cols);
    const endTime = performance.now();
    console.log((endTime - startTime) + 'ms'); // 何ミリ秒かかったかを表示

    return pred
}

setTimeout(() => {
    const input = [1.247e+01, 1.860e+01, 8.109e+01, 4.819e+02, 9.965e-02, 1.058e-01, 8.005e-02, 
                   3.821e-02, 1.925e-01, 6.373e-02, 3.961e-01, 1.044e+00, 2.497e+00, 3.029e+01,
                   6.953e-03, 1.911e-02, 2.701e-02, 1.037e-02, 1.782e-02, 3.586e-03, 1.497e+01,
                   2.464e+01, 9.605e+01, 6.779e+02, 1.426e-01, 2.378e-01, 2.671e-01, 1.015e-01,
                   3.014e-01, 8.750e-02];  // 1次元配列で定義
    const rows = 1;  // データ数
    const cols = 30;  // 特徴量数
    // 期待出力結果: [0.99957539]

    console.log('Input data: ', input);
    let output = model_inference(input, rows, cols);
    console.log('Prediction result: ', output)

}, 1000);

window.onload = loadWasmModule;
