function isDarkMode() {
    return window.matchMedia('(prefers-color-scheme: dark)').matches;
}

// Initialize Monaco Editor
require.config({
    paths: { vs: 'assets/library/node_modules/monaco-editor/min/vs' }
});

require(['vs/editor/editor.main'], function () {

    // Initialize editors
    let eqlEditor = monaco.editor.create(document.getElementById('eql-editor'), {
        value: ['entities.select("_entity")'].join('\n'),
        language: 'javascript',
        theme: 'vs', // Default theme
        automaticLayout: true
    });

    window.eqlEditor = eqlEditor;

    // Initialize HERY Monaco editor
    let heryEditor = monaco.editor.create(document.getElementById('hery-editor'), {
        //value: ['_entity: github.com/AmadlaOrg/Entity', '_body:', '\tname: "John Doe"'].join('\n'),
        language: 'yaml',
        theme: 'vs', // Default theme
        automaticLayout: true
    });

    window.heryEditor = heryEditor;

    // JSON object we want to edit
    const jsonCode = [{
        "enabled": true,
        "description": "something"
    }];

    const modelUri = monaco.Uri.parse("json://grid/settings.json");
    const jsonModel = monaco.editor.createModel(JSON.stringify(jsonCode, null, '\t'), "json", modelUri);

    monaco.editor.create(document.getElementById('output-editor'), {
        readOnly: true,
        model: jsonModel,
        theme: 'vs', // Default theme
        automaticLayout: true
    });

    const isDark = document.documentElement.classList.contains('pf-v6-theme-dark');
    monaco.editor.setTheme(isDark ? 'vs-dark' : 'vs');
});

function theme() {
    return {
        theme: 'light', // Default tab

        init() {
            if (isDarkMode()) {
                this.setTheme('dark');
            } else {
                this.setTheme(this.theme)
            }
        },

        setTheme(theme) {
            this.theme = theme;

            // Update Monaco Editor theme
            if (typeof monaco === "object") {
                const monacoTheme = this.theme === 'dark' ? 'vs-dark' : 'vs';
                monaco.editor.setTheme(monacoTheme);
            }
        }
    };
}

/*function filesTabs() {
    return {
        files: [],
        readFile(event) {
            const file = event.target.files[0];
            if (!file) return;

            const reader = new FileReader();
            reader.onload = (e) => {
                const fileContent = e.target.result;
                console.log("File content:", fileContent);


                const cleanName = file.name.replace(/\./g, '-');
                this.files.push({
                    id: Date.now(), // unique ID for x-for key
                    name: file.name,
                    cleanName: cleanName,
                    content: fileContent
                });
            };
            reader.readAsText(file);
        }
    };
}*/

function filesTabs() {
    return {
        files: [],
        readFile(event) {
            const file = event.target.files[0];
            if (!file) return;

            const reader = new FileReader();
            reader.onload = (e) => {
                const fileContent = e.target.result;
                const cleanName = file.name.replace(/\./g, '-');
                this.files.push({
                    id: Date.now(),
                    name: file.name,
                    cleanName: cleanName,
                    content: fileContent
                });

                // Set the content into the HERY editor
                if (window.heryEditor) {
                    window.heryEditor.setValue(fileContent);
                } else {
                    console.warn("HERY editor is not initialized yet.");
                }
            };
            reader.readAsText(file);
        }
    };
}
