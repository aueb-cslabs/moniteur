import React from 'react';
import dateformat from 'dateformat';
import './TitleBar.scss';

class TitleBar extends React.Component {

    constructor(props) {
        super(props);

        this.state = {date: new Date()};

        setInterval(function () {
            this.setState( {date: new Date()})
        }.bind(this), 1000);
    }

    style() {
        return {
            'background-color': '#762124'
        }
    }

    time() {
        return dateformat(this.state.date, 'HH:MM:ss')
    }

    render() {
        return (
            <nav className="navbar navbar-expand navbar-dark" style={this.style()}>

                <a className='navbar-brand'>
                    <img src={'https://www.aueb.gr/press/logos/1_AUEB-pantone-HR.jpg'}
                         height="44" />
                </a>
                {/*<a className="navbar-brand">AUEB CSLabs</a>*/}

                <ul className="navbar-nav mr-auto">
                    <li className="nav-item active">
                        <a className="nav-link" href="#">Home <span className="sr-only">(current)</span></a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="#">Link</a>
                    </li>
                </ul>
                <a className="navbar-brand">
                    {this.time()}
                </a>
            </nav>
        );
    }
}

export default TitleBar;
