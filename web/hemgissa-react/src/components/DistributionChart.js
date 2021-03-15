import "../assets/css/blk-design-system.min.css";
import "../App.css";
import {Line} from 'react-chartjs-2';
import * as React from "react";

class DistributionChart extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            name: this.props.name,
            chart: {
                labels: this.props.lables,
                datasets: [
                    {
                        label: 'Distribution',
                        fill: true,
                        lineTension: 0.1,
                        backgroundColor: 'rgba(61,162,162,0.4)',
                        borderColor: 'rgb(75,192,192)',
                        borderCapStyle: 'butt',
                        borderDash: [],
                        borderDashOffset: 0.0,
                        borderJoinStyle: 'miter',
                        pointBorderColor: 'rgba(75,192,192,1)',
                        pointBackgroundColor: '#fff',
                        pointBorderWidth: 1,
                        pointHoverRadius: 5,
                        pointHoverBackgroundColor: 'rgba(75,192,192,1)',
                        pointHoverBorderColor: 'rgba(220,220,220,1)',
                        pointHoverBorderWidth: 2,
                        pointRadius: 1,
                        pointHitRadius: 10,
                        data: this.props.data
                    }
                ]
            }
        }
    };


    render() {
        return (
            <div>
                <h3>{this.state.name} Distribution Plot</h3>
                <Line options={{maintainAspectRatio: true}} height={200} data={this.state.chart}/>
            </div>
        );
    }
}

export default DistributionChart;