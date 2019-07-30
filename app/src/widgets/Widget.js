import React from 'react';

export default class Widget extends React.Component {

    classes() {
        let width = this.props.width;
        let height = this.props.height;
        return `w-${width * 25} h-${height * 25}`
    }

    render() {
        return <div className={"Widget p-4 " + this.classes()}>
            <div className="card h-100 w-100">
                <div className={"card-body h-100 w-100 d-flex justify-content-center align-items-center " + this.props.className}>
                    {this.props.children}
                </div>
            </div>
        </div>
    }

}