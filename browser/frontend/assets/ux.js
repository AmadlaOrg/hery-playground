/*function isDarkMode() {
    return window.matchMedia('(prefers-color-scheme: dark)').matches;
}

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
        }
    };
}*/

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
