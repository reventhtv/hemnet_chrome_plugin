const chartCollectRules = {
    rent: {
        searchMin: 0,
        searchMax: 500,
        interval: 500,
    },
    size: {
        searchMin: 0,
        searchMax: 5,
        interval: 5,
    },
    rooms: {
        searchMin: 0,
        searchMax: 1,
        interval: 1,
    },
}

export function getChartData(values, name, globalMax) {

    let graph = new Map();
    values.sort((a, b) => a - b);
    graph.set(0, 0);

    let rules;
    //Rent,Rooms, Size
    if (name === "Rent") {
        rules = chartCollectRules.rent;
    }
    if (name === "Rooms") {
        rules = chartCollectRules.rooms;
    }
    if (name === "Size") {
        rules = chartCollectRules.size;
    }

    let localSearchMin = rules.searchMin;
    let localSearchMax = rules.searchMax;
    let localInterval = rules.interval;

    values.forEach(function (v) {
        while (localSearchMin < globalMax) {
            if (v > localSearchMin && v <= localSearchMax) {
                if (graph.has(localSearchMax)) {
                    graph.set(localSearchMax, graph.get(localSearchMax) + 1)
                } else {
                    graph.set(localSearchMax, 1)
                }
                break;
            } else {
                localSearchMin = localSearchMax;
                localSearchMax = localSearchMin + localInterval;
            }
        }
    });
    let labels = Array.from(graph.keys())
    let data = Array.from(graph.values())
    return {chartLabels: labels, chartValues: data};
}
