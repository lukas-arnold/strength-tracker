Chart.defaults.elements.line.tension = 0.4;
Chart.defaults.interaction = {
    mode: "index",
    intersect: false
};

function renderChart(canvasId, model) {
    const ctx = document.getElementById(canvasId);

    new Chart(ctx, {
        type: "line",
        data: {
            labels: model.labels,
            datasets: [
                {
                    label: model.loadLabel,
                    data: model.loadData,
                    yAxisID: "y"
                },
                {
                    label: model.repetitionsLabel,
                    data: model.repetitionsData,
                    yAxisID: "y1"
                }
            ]
        },
        options: {
            plugins: {
                tooltip: {
                    callbacks: {
                        title: (items) => formatDate(items[0].label),
                        label: (ctx) => {
                            const formatter =
                                ctx.dataset.yAxisID === "y"
                                    ? model.loadFormatter
                                    : value => value;

                            return `${ctx.dataset.label}: ${formatter(ctx.raw)}`;
                        }
                    }
                }
            },

            scales: {
                y: {
                    position: "left",
                    ticks: {
                        callback: (value) =>
                            model.loadFormatter(value)
                    }
                },

                y1: {
                    position: "right",
                    beginAtZero: true,
                    ticks: {
                        stepSize: 1
                    }
                },

                x: {
                    ticks: {
                        callback(value) {
                            return formatDate(
                                this.getLabelForValue(value)
                            );
                        }
                    }
                }
            }
        }
    });
}