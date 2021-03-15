import "../assets/css/blk-design-system.min.css";
import "../assets/css/nucleo-icons.css"
import * as React from "react";

class Banner extends React.Component {
    constructor(props) {
        super(props);
        const oneMill = 1000000
        this.state = {
            name: this.props.entityProps.name,
            region: this.props.entityProps.region,
            imageUrl: this.props.entityProps.imageUrl,
            askingPrice: (this.props.entityProps.askingPrice / oneMill),
            estimatedPrice: (this.props.entityProps.estimatedValue / oneMill),
        };
    }

    render() {
        return (
            <div className="row row-grid justify-content-between">
                <div className="col-md-6 mt-lg-5">
                    <div className="carousel">
                        <div className="carousel-item active">
                            <img className="d-block w-100"
                                 src={this.state.imageUrl}
                                 alt="First slide"/>
                        </div>
                    </div>
                    <div className="row">
                        <div className="col-lg-6 col-sm-12 px-2 py-2">
                            <div className="card card-stats ">
                                <div className="card-body">
                                    <div className="row">
                                        <div className="col-5 col-md-4">
                                            <div className="icon-big text-center icon-warning">
                                                <i className="tim-icons icon-trophy text-warning"/>
                                            </div>
                                        </div>
                                        <div className="col-7 col-md-8">
                                            <div className="numbers">
                                                <p className="card-title">{this.state.askingPrice} M
                                                </p>
                                                <p>
                                                </p>
                                                <p className="card-category">Asking Price</p>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div className="col-lg-6 col-sm-12 px-2 py-2">
                            <div className="card card-stats upper bg-default">
                                <div className="card-body">
                                    <div className="row">
                                        <div className="col-5 col-md-4">
                                            <div className="icon-big text-center icon-warning">
                                                <i className="tim-icons icon-coins text-white"/>
                                            </div>
                                        </div>
                                        <div className="col-7 col-md-8">
                                            <div className="numbers">
                                                <p className="card-title">{this.state.estimatedPrice} M
                                                </p>
                                                <p>
                                                </p>
                                                <p className="card-category">Estimated</p>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="col-md-6 mt-lg-5">
                    <div className="pl-md-5">
                        <h1>{this.state.name}</h1>
                        <h2>{this.state.region}</h2>
                        <p>I should be capable of drawing a single stroke at the present moment; and yet I feel that
                            I
                            never
                            was a greater artist than now.

                            When, while the lovely valley teems with vapour around me, and the meridian sun strikes
                            the
                            upper
                            surface of the impenetrable foliage of my trees, and but a few stray.</p>
                    </div>
                </div>
            </div>
        );
    }
}

export default Banner;