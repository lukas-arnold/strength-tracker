
Chart.defaults.elements.line.tension = 0.4;

function renderChart(canvasId, model) {
    const ctx = document.getElementById(canvasId);

    new Chart(ctx, {
        type: model.type || "line",
        data: {
            labels: model.labels,
            datasets: [{
                label: model.label,
                data: model.data,
                borderWidth: model.borderWidth || 1
            }]
        },
        options: {
            plugins: {
                tooltip: {
                    callbacks: {
                        title: (items) => formatDate(items[0].label),
                        label: (ctx) => formatLoad(ctx.formattedValue)
                    }
                }
            },
            scales: {
                y: {
                    beginAtZero: false,
                    ticks: {
                        callback(value) {
                            return formatLoad(value);
                        }
                    }
                },
                x: {
                    ticks: {
                        callback(value) {
                            return formatDate(this.getLabelForValue(value));
                        }
                    }
                }
            }
        }
    });
}
