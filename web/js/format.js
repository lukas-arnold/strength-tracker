document.addEventListener('DOMContentLoaded', function() {
    formatDates();
    formatLoads();
});

function formatDates() {
    const dateElements = document.querySelectorAll(".date");
    
    dateElements.forEach(element => {
        const originalDate = new Date(element.textContent);
        const options = {year: "numeric", month: "2-digit", day: "2-digit"};
        element.textContent = originalDate.toLocaleDateString(undefined, options);
    });
}

function formatLoads() {
    const loadElements = document.querySelectorAll(".load");

    loadElements.forEach(element => {
        let load = element.textContent;
        load = load.replace(".", ",");
        load += " kg";
        element.textContent = load
    });
}