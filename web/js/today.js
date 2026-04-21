document.addEventListener('DOMContentLoaded', function() {
    const dateField = document.getElementById("date");
    setCurrentDate(dateField);
});

const setCurrentDate = (dateField) => {
    const today = new Date();
    const offset = today.getTimezoneOffset();
    const todayUTC = new Date(today.getTime() + (offset*60*1000));
    const currentDate = todayUTC.toISOString().split('T')[0];
    dateField.value = currentDate;
}