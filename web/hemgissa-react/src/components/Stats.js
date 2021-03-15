import "../assets/css/blk-design-system.min.css";
import * as React from "react";
import Banner from "./Banner";
import Summary from "./Summary";

class Stats extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            entityProps: this.props.response.entity,
            countySummary: {
                name: this.props.response.countySummary.name,
                size: this.props.response.countySummary.size,
                rent: this.props.response.countySummary.rent,
                rooms: this.props.response.countySummary.rooms,
            }
        };
    }


    render() {
        return (
            <>
                <div className="container">
                    <Banner entityProps={this.state.entityProps}/>
                </div>
                <div className="container">
                    <div className="title">
                        <h3 className="mb-3">{this.state.countySummary.name} Stats</h3>
                    </div>
                    <Summary name="Size" values={this.state.countySummary.size}/>
                    <Summary name="Rent" values={this.state.countySummary.rent}/>
                    <Summary name="Rooms" values={this.state.countySummary.rooms}/>
                </div>

                <div className="container">
                    <div className="title">
                        <h3 className="mb-3">{this.state.countySummary.name} Stats</h3>
                    </div>
                    <Summary name="Size" values={this.state.countySummary.size}/>
                    <Summary name="Rent" values={this.state.countySummary.rent}/>
                    <Summary name="Rooms" values={this.state.countySummary.rooms}/>
                </div>
            </>

        );
    }
}

export default Stats;