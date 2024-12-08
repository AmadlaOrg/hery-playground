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


