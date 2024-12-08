// Initialize Monaco Editor
/*require.config({ paths: { 'vs': 'assets/library/node_modules/monaco-editor/min/vs' }});
require(['vs/editor/editor.main'], function() {
    var editor = monaco.editor.create(document.getElementById('monaco-container'), {
        value: "// Monaco Editor is ready!",
        language: 'javascript',
        theme: 'vs-dark',
        automaticLayout: true
    });

    // Ensure Monaco resizes to fill its container
    editor.layout();
});*/
/*require.config({ paths: { vs: 'assets/library/node_modules/monaco-editor/min/vs' } });
require(['vs/editor/editor.main'], function () {
    var eqlEditor = monaco.editor.create(document.getElementById('eql-editor'), {
        value: ['function x() {', '\tconsole.log("Hello world!");', '}'].join('\n'),
        language: 'javascript'
    });
    var heryEditor = monaco.editor.create(document.getElementById('hery-editor'), {
        value: ['function x() {', '\tconsole.log("Hello world!");', '}'].join('\n'),
        language: 'javascript'
    });
});*/
// Monaco Editor configuration
/*require.config({ paths: { vs: 'assets/library/node_modules/monaco-editor/min/vs' } });

// Wait for Monaco to be loaded, then initialize editors
require(['vs/editor/editor.main'], function() {
    var eqlEditor = monaco.editor.create(document.getElementById('eql-editor'), {
        value: ['function x() {', '\tconsole.log("Hello world!");', '}'].join('\n'),
        language: 'javascript',
        theme: 'vs-dark',
        automaticLayout: true
    });

    var heryEditor = monaco.editor.create(document.getElementById('hery-editor'), {
        value: ['function x() {', '\tconsole.log("Hello world!");', '}'].join('\n'),
        language: 'javascript',
        theme: 'vs-dark',
        automaticLayout: true
    });
});*/
require.config({
    paths: { vs: 'assets/library/node_modules/monaco-editor/min/vs' }
});

require(['vs/editor/editor.main'], function () {
    // Initialize EQL Monaco editor
    let eqlEditor = monaco.editor.create(document.getElementById('eql-editor'), {
        value: ['entities.select("_entity")'].join('\n'),
        language: 'javascript',
        //theme: 'vs-dark',
        automaticLayout: true
    });

    // Initialize HERY Monaco editor
    let heryEditor = monaco.editor.create(document.getElementById('hery-editor'), {
        value: ['_entity: github.com/AmadlaOrg/Entity', '_body:', '\tname: "John Doe"'].join('\n'),
        language: 'yaml',
        //theme: 'vs-dark',
        automaticLayout: true
    });

    /*let editor = monaco.editor.create(document.getElementById('output-editor'), {
        value: '{\n"message": "Hello World!", "internal": {"name": "John Doe"}\n}',
        language: 'json',
        readOnly: true, // Set the editor to read-only
        automaticLayout: true, // Automatically adjust layout for the editor
        //theme: 'vs-dark' // Use a dark theme (optional)
    });

    // Optional: Format the JSON data (to ensure it's nicely indented)
    const model = editor.getModel();
    monaco.editor.format(model);*/

    // JSON object we want to edit
    const jsonCode = [{
        "enabled": true,
        "description": "something"
    }];

    const modelUri = monaco.Uri.parse("json://grid/settings.json");
    const jsonModel = monaco.editor.createModel(JSON.stringify(jsonCode, null, '\t'), "json", modelUri);

    const editor = monaco.editor.create(document.getElementById('output-editor'), {
        readOnly: true,
        model: jsonModel
    });
});
