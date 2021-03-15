import "../assets/css/blk-design-system.min.css";
import * as React from "react";
import homePicture from "../assets/img/denys.jpg";
import testData from "./test_data/response.json"
import Stats from "./Stats";

class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            guessUrl: 'test',
            renderStats: false,
            response: {}
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleClick = this.handleClick.bind(this);
    }

    handleChange(e) {
        this.setState({guessUrl: e.target.value});
    }

    handleClick(e) {
        if (this.state.guessUrl || this.state.guessUrl !== "") {
            this.setState({
                renderStats: true,
                response: testData
            });
        }

        e.preventDefault();
    }


    render() {
        if (!this.state.renderStats) {
            return (
                <div className="container">
                    <div>
                        <div className="title">
                            <h2>Hemgissa</h2>
                        </div>
                        <div className="row justify-content-between align-items-center">
                            <div className="col-lg-5 mb-5 mb-lg-0">
                                <h2 className="text-white font-weight-light">Get price estimate and stats for apartments
                                    in
                                    Stockholm</h2>
                                <p className="text-white mt-4">
                                    Black Design comes with three pre-built pages
                                    to help you get started faster. You can change the text and images and you're good
                                    to
                                    go.</p>

                                <div className="form-group">
                                    <input type="text" value={this.state.guessUrl} placeholder="Hemnet Link"
                                           className="form-control" onChange={this.handleChange}/>
                                </div>
                                <button className="btn btn-primary" onClick={this.handleClick}>
                                    Gissa
                                </button>
                            </div>
                            <div className="col-lg-6">
                                <div id="carouselExampleControls" className="carousel slide">
                                    <div className="carousel-inner">
                                        <div className="carousel-item active">
                                            <img className="d-block w-100" src={homePicture} alt="First slide"/>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            );
        } else {
            return (
                <Stats response={this.state.response}/>
            )
        }
    }
}

export default Home;