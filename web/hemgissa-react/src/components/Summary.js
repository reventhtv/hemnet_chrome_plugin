import "../assets/css/blk-design-system.min.css";
import "../assets/css/nucleo-icons.css"
import * as React from "react";
import DistributionChart from "./DistributionChart";

class Summary extends React.Component {

    setChartData(vals) {
        let graph = new Map()
        vals.sort(function (a, b) {
            return a - b;
        });
        graph.set(0, 0)
        vals.forEach(function (v) {
            if (graph.has(v)) {
                graph.set(v, graph.get(v) + 1)
            } else {
                graph.set(v, 1)
            }
        });
        let labels = Array.from(graph.keys())
        let values = Array.from(graph.values())
        return {chartLabels: labels, chartValues: values};
    }


    constructor(props) {
        super(props);
        let graph = this.setChartData(this.props.values.values)
        this.state = {
            name: this.props.name,
            min: this.props.values.min,
            max: this.props.values.max,
            mean: this.props.values.mean,
            median: this.props.values.median,
            values: this.props.values.values,
            chartLabels: graph.chartLabels,
            chartValues: graph.chartValues,
        };
        this.setChartData(this.props.values.values)
    }

    render() {
        return (
            <div className="row">
                <div className="col-md-10 ml-auto col-xl-6 mr-auto">
                    <div className="card">
                        <div className="card-header">
                            <ul className="nav nav-tabs nav-tabs-primary" role="tablist">
                                <li className="nav-item">
                                    <div className="nav-link active">
                                        {this.state.name}
                                    </div>
                                </li>
                            </ul>
                        </div>
                        <div className="card-body">
                            <div className="tab-content tab-subcategories">
                                <div className="tab-pane active" id="linka">
                                    <div className="table-responsive">
                                        <table className="table tablesorter" id="plain-table">
                                            <thead className=" text-primary">
                                            <tr>
                                                <th className="header">
                                                </th>
                                                <th className="header">
                                                    VALUE
                                                </th>
                                            </tr>
                                            </thead>
                                            <tbody>
                                            <tr>
                                                <td>
                                                    MIN
                                                </td>
                                                <td>
                                                    {this.state.min}
                                                </td>
                                            </tr>
                                            <tr>
                                                <td>
                                                    MAX
                                                </td>
                                                <td>
                                                    {this.state.max}
                                                </td>
                                            </tr>
                                            <tr>
                                                <td>
                                                    MEAN
                                                </td>
                                                <td>
                                                    {this.state.mean}
                                                </td>
                                            </tr>
                                            <tr>
                                                <td>
                                                    MEDIAN
                                                </td>
                                                <td>
                                                    {this.state.median}
                                                </td>
                                            </tr>
                                            <tr>
                                                <td>
                                                    CURRENT
                                                </td>
                                                <td>
                                                    {this.state.max}
                                                </td>
                                            </tr>
                                            </tbody>
                                        </table>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="col-md-10 ml-auto col-xl-6 mr-auto">
                    <DistributionChart name={this.state.name} lables={this.state.chartLabels}
                                       data={this.state.chartValues}/>
                </div>
            </div>
        );
    }
}

export default Summary;