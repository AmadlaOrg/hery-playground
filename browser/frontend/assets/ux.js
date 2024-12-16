/**
 * All about the different components inside the page that are not related to Monaco editor.
 *
 * Some of the components need to have access to the Monaco editor elements so they are in the editor.js file.
 */

/**
 *
 * @returns {{init(): void, setTab(*): void, activeTab: string}}
 */
function tabSystem() {
    return {
        activeTab: 'hery', // Default tab

        init() {
            this.setTab(this.activeTab);
        },

        setTab(tab) {
            this.activeTab = tab;
        }
    };
}

/**
 *
 * @returns {{wasmSupported: boolean, fileApiSupported: boolean, fetchSupported: boolean}}
 */
function featureCheck() {
    return {
        wasmSupported: (typeof WebAssembly === "object" && typeof WebAssembly.instantiate === "function"),
        fileApiSupported: (typeof File !== 'undefined' && typeof FileReader !== 'undefined' && typeof Blob !== 'undefined'),
        fetchSupported: (typeof fetch === 'function')
    }
}

/**
 *
 * @returns {{init(): void, setInfoModal(*): void, infoModalShow: boolean}}
 */
function page() {
    return {
        infoModalShow: false,
        settingsModalShow: false,
        newWorkspaceModalShow: false,

        init() {
            this.setInfoModal(this.infoModal)
            this.setSettingsModal(this.settingsModalShow)
            this.setNewWorkspaceModalShow(this.newWorkspaceModalShow)
        },

        setInfoModal(status) {
            this.infoModal = status;
        },

        setSettingsModal(status) {
            this.settingsModalShow = status;
        },

        setNewWorkspaceModalShow(status) {
            this.newWorkspaceModalShow = status
        },
    };
}

/*function readFile(event) {
    console.log("readFile")
    const file = event.target.files[0];
    if (!file) return;
    const reader = new FileReader();
    reader.onload = (e) => {
        this.fileContent = e.target.result;
        console.log("File content:", this.fileContent);
    };
    reader.readAsText(file);
    console.log(file)

    console.log(file.name)
}*/

/*function filesTabs() {
    return {
        files: []
    }
}

function readFile(event) {
    const file = event.target.files[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = (e) => {
        const fileContent = e.target.result;
        console.log("File content:", fileContent);

        const cleanName = file.name.replace(/\./g, '-');
        this.files.push({
            id: Date.now(), // unique ID for the v-for key
            name: file.name,
            cleanName: cleanName,
            content: fileContent
        });
    };
    reader.readAsText(file);
}*/


