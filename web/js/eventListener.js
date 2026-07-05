document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll(".date")
        .forEach(el => el.textContent = formatDate(el.textContent));

    document.querySelectorAll(".load")
        .forEach(el => el.textContent = formatLoad(el.textContent));

    const currentDate = new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .split('T')[0];

    document.querySelectorAll('input[type="date"]').forEach(input => {
        if (!input.value) {
            input.value = currentDate;
        }
    });
});
