function tabSystem() {
    return {
        activeTab: 'hery', // Default tab

        init() {
            this.setTab(this.activeTab);
        },

        setTab(tab) {
            this.activeTab = tab;

            // Move the content from the hidden tab to the main content area
            //const tabEls = document.getElementsByClassName(`tab_${tab}`);
        }
    };
}
