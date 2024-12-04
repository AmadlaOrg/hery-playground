require.config({
    paths: { 'vs': 'https://cdn.jsdelivr.net/npm/monaco-editor@0.34.1/min/vs' }
});

require(['vs/editor/editor.main'], function() {
    // Register custom EQL language
    monaco.languages.register({ id: 'eql' });

    // Define syntax for EQL
    monaco.languages.setMonarchTokensProvider('eql', {
        tokenizer: {
            root: [
                [/\b(entities|select|meta|body|equal)\b/, 'keyword'],
                [/"([^"]*)"/, 'string'],
                [/\b([a-zA-Z_][a-zA-Z0-9_]*)\b/, 'identifier'],
            ]
        }
    });

    // Add autocomplete suggestions for EQL
    monaco.languages.registerCompletionItemProvider('eql', {
        provideCompletionItems: function() {
            var suggestions = [
                { label: 'select', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'select' },
                { label: 'equal', kind: monaco.languages.CompletionItemKind.Keyword, insertText: 'equal' },
                { label: 'id', kind: monaco.languages.CompletionItemKind.Text, insertText: 'id' },
                { label: '_entity', kind: monaco.languages.CompletionItemKind.Text, insertText: '_entity' }
            ];
            return { suggestions: suggestions };
        }
    });

    // Create the Monaco editor with custom EQL language
    var editor = monaco.editor.create(document.getElementById('container'), {
        value: `entities\nentities.select("id")\nentities.select("id", "_entity")\nmeta.select("id")\nbody.equal("id", "some-id-to-equal")`,
        language: 'eql'
    });
});
