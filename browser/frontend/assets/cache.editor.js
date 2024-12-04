let go = new Go(); // Defined in wasm_exec.js
let db;

// Load Go WebAssembly binary
WebAssembly.instantiateStreaming(fetch('browser.wasm'), go.importObject).then(result => {
    go.run(result.instance);

    // Interact with the WebAssembly instance, e.g., call the 'add' function
    window.add(3, 4); // Call Go's 'add' function
});

// Handle drag and drop
document.getElementById('dropzone').addEventListener('drop', function(e) {
    e.preventDefault();
    e.stopPropagation();
    const file = e.dataTransfer.files[0];
    const reader = new FileReader();

    reader.onload = function(event) {
        // Open SQLite database from the dropped file
        SQL.load({locateFile: file => 'https://cdnjs.cloudflare.com/ajax/libs/sql.js/1.12.0/'})
            .then(SQL => {
                db = new SQL.Database(new Uint8Array(event.target.result));
                document.getElementById('output').textContent = "Database opened successfully!";
                displayDatabase();
            });
    };

    reader.readAsArrayBuffer(file);
});

// Display table names in the SQLite database
function displayDatabase() {
    try {
        const res = db.exec("SELECT name FROM sqlite_master WHERE type='table';");
        let output = 'Tables in database:\n';
        res[0].values.forEach(table => {
            output += table[0] + '\n';
        });
        document.getElementById('output').textContent = output;
    } catch (err) {
        console.error(err);
        document.getElementById('output').textContent = 'Error loading database.';
    }
}
