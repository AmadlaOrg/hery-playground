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
/*function page() {
    return {
        // Modals
        infoModalShow: false,
        settingsModalShow: false,
        newWorkspaceModalShow: false,

        // Elements size
        headerHeight: 0,
        headerToolbarOptionsWidth: 0,

        // Document width and height that changes with resizing
        documentWidth: 0,
        documentHeight: 0,

        init() {
            this.setInfoModal(this.infoModal)
            this.setSettingsModal(this.settingsModalShow)
            this.setNewWorkspaceModalShow(this.newWorkspaceModalShow)
            this.setHeaderHeight()
            this.setHeaderToolbarOptionsWidth()
            this.updateDocumentSize(); // Initialize size on page load
            window.addEventListener('resize', this.updateDocumentSize.bind(this));
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

        setHeaderHeight() {
            const header = document.getElementById('header');
            if (header) {
                this.headerHeight = header.getBoundingClientRect().height;
            }
        },

        setHeaderToolbarOptionsWidth() {
            const toolbarOptions = document.getElementById('header-toolbar-options');
            if (toolbarOptions) {
                this.headerToolbarOptionsWidth = toolbarOptions.getBoundingClientRect().width;
            }
        },

        updateDocumentSize() {
            this.documentWidth = window.innerWidth;
            this.documentHeight = window.innerHeight;
        },

        calculateMenuPosition() {
            // Use dynamic header height and add 25px
            const menuTop = this.headerHeight + 25;

            // Position menu relative to the toolbar options
            const menuLeft = this.headerToolbarOptionsWidth;

            return `transform: translate(${menuLeft}px, ${menuTop}px);`;
        },

        // :style="position: absolute; inset: 0px auto auto 0px;
        // ${calculateMenuPosition()};
        // width: 246px;
        // z-index: 9999;
        // opacity: 1;
        // transition: opacity cubic-bezier(0.54, 1.5, 0.38, 1.11);"
        get menuStyle() {
            return {
                position: 'absolute',
                inset: '0px auto auto 0px',
                ...this.calculateMenuPosition(),
                width: '246px',
                zIndex: '9999',
                opacity: '1',
                transition: 'opacity cubic-bezier(0.54, 1.5, 0.38, 1.11)',
            };
        },
    };
}*/

function page() {
    return {
        // Modals
        infoModalShow: false,
        settingsModalShow: false,
        newWorkspaceModalShow: false,

        // Elements size
        headerHeight: 0,
        headerToolbarOptionsWidth: 0,

        // Document width and height that changes with resizing
        documentWidth: 0,
        documentHeight: 0,

        // Menu style object
        menuStyle: {},

        // Workspace menu
        searchWorkspace: '', // Current search input value
        // TODO: this needs to be removed
        workspaceMenuItems: [
            "A very, very long name to test what happens",
            "Action 2 - aaa",
            "Action 3 - longer!",
            "Action 4 - making longer and longer and longer",
            "Action 5 - longer",
            "Action 6 - Change again",
            "Action 7",
            "Action 8",
            "Action 9",
            "Action 10",
            "Action 11",
            "Action 12"
        ],
        workspaceFilteredItems: [],

        workspaceFilterMenu() {
            // Filter items based on the search term
            const term = this.searchWorkspace.toLowerCase();
            this.workspaceFilteredItems = this.workspaceMenuItems.filter(item =>
                item.toLowerCase().includes(term)
            );
        },

        init() {
            this.setInfoModal(this.infoModalShow);
            this.setSettingsModal(this.settingsModalShow);
            this.setNewWorkspaceModalShow(this.newWorkspaceModalShow);
            this.updateDocumentSize(); // Initialize size on page load
            this.setHeaderHeight();
            this.setHeaderToolbarOptionsWidth();
            this.updateMenuStyle(); // Initialize menu style
            this.workspaceFilteredItems = this.workspaceMenuItems;
            window.addEventListener('resize', this.handleResize.bind(this));
        },

        setInfoModal(status) {
            this.infoModalShow = status;
        },

        setSettingsModal(status) {
            this.settingsModalShow = status;
        },

        setNewWorkspaceModalShow(status) {
            this.newWorkspaceModalShow = status;
        },

        setHeaderHeight() {
            const header = document.getElementById('header');
            if (header) {
                this.headerHeight = header.getBoundingClientRect().height;
            }
        },

        setHeaderToolbarOptionsWidth() {
            const toolbarOptions = document.getElementById('header-toolbar-options');
            if (toolbarOptions) {
                this.headerToolbarOptionsWidth = toolbarOptions.getBoundingClientRect().width;
            } else {
                console.warn("Toolbar options element not found.");
                this.headerToolbarOptionsWidth = 0; // Fallback to avoid NaN
            }
        },

        updateDocumentSize() {
            this.documentWidth = window.innerWidth;
            this.documentHeight = window.innerHeight;
        },

        calculateMenuPosition() {
            // Debug log
            /*console.log("Header Toolbar Options Width:", this.headerToolbarOptionsWidth);
            console.log("Header Height:", this.headerHeight);

            const menuTop = this.headerHeight + 25; // Dynamic vertical position
            const menuLeft = this.headerToolbarOptionsWidth; // Dynamic horizontal position

            return `translate(${menuLeft}px, ${menuTop}px)`;*/
            // Use dynamic document width and calculate menu left position
            const menuTop = this.headerHeight + 4; // Add header height and offset
            const dropdownMenu = document.getElementById('dropdown-menu-workspace');
            let menuLeft = 0;

            if (dropdownMenu) {
                //console.log("this.headerToolbarOptionsWidth: --- "+this.headerToolbarOptionsWidth)
                //console.log("dropdownMenu.getBoundingClientRect().width: "+dropdownMenu.getBoundingClientRect().width)
                const dropdownWidth = this.headerToolbarOptionsWidth; //dropdownMenu.getBoundingClientRect().width;
                menuLeft = (this.documentWidth - dropdownWidth) - 4; // Align to document width
            } else {
                console.warn("Dropdown menu element not found.");
            }

            /*console.log("Document Width:", this.documentWidth);
            console.log("Menu Top:", menuTop);
            console.log("Menu Left:", menuLeft);*/

            return `translate(${menuLeft}px, ${menuTop}px)`;
        },

        updateMenuStyle() {
            this.menuStyle = {
                position: "absolute",
                inset: "0px auto auto 0px",
                width: "246px",
                zIndex: "9999",
                opacity: "1",
                transition: "opacity cubic-bezier(0.54, 1.5, 0.38, 1.11)",
                transform: this.calculateMenuPosition(),
            };
        },

        handleResize() {
            this.updateDocumentSize();
            this.setHeaderHeight();
            this.setHeaderToolbarOptionsWidth();
            this.updateMenuStyle(); // Recalculate menu style on resize
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


