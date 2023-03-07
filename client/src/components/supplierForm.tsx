import React, {Component} from 'react';  
import axios from "axios";
import {Card, Header, Form, Input, Icon} from "semantic-ui-react";

let endpoint = "http://localhost:9000"

interface IProps {}
interface IState {
    supplier: string
    address: string
    logo: string 
    items: []
}

class Supplier extends Component<IProps, IState> {
    constructor(props: any) {
        super(props);
        this.state = {
            supplier: "",
            address: "",
            logo: "",
            items:[],
        };
    }
    componentDidMount() {
        this.getSupplier();
    }

    getSupplier = () => {
        axios.get(endpoint + "/api/supplier") .then((res) => {
            if(res.data) {
                this.setState({
                    items: res.data.map((item: any) => {
                    return (
                        <Card key={item._id}  fluid className="rough">
                            <Card.Content>
                            <Card.Meta textAlign="left">
                                <Icon
                                name="delete"
                                color="blue"
                                onClick={()=> this.deleteSupplier(item._id)}
                                />
                                <span style={{paddingRight: 10}}>Delete</span>
                            </Card.Meta>
                                <Card.Header textAlign="center">
                                    {item.supplier}
                                </Card.Header>
                                <Card.Description textAlign="center">
                                    <div style={{wordWrap: "break-word"}}>{item.address}</div>
                                    <img src={item.logo}/>
                                </Card.Description>
                            </Card.Content>

                        </Card>
                    )
                    }
                )})
            } else {
                this.setState({
                    items:[],
                });
            }
        })
    }

    onSubmit = () => {
        let supplier = this.state.supplier;
        let address = this.state.address;
        let logo = this.state.logo;
        if (supplier && address && logo) {
            axios.post(endpoint + "/api/supplier", {supplier, address, logo}, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
            }).then((res)=> {
                this.getSupplier();
                this.setState ({
                    supplier:"",
                    address: "",
                    logo: "",
                });
                console.log(res);
            });
        }
    }


    updateSupplier = (id:string) => {
        axios.put(endpoint +"/api/supplier/" + id, {
            headers: {
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=> {
            console.log(res);
            this.getSupplier();
        })
    }

    deleteSupplier = (id: string) => {
        axios.delete(endpoint + "/api/deleteSupplier/" + id,{
            headers: {
                "Content-Type": "application:/x-www-form-urlencoded",
            },
        }).then((res) => {
            this.getSupplier();
        })
    }

    onChange = (e: React.FormEvent<HTMLInputElement>) => {
        this.setState({
            [e.currentTarget.name] : e.currentTarget.value
        } as any);
    }

   //TODO: write reject
    encodeBase64 = (file:File) => {
        return new Promise((resolve => {
            const reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = () => {
                resolve(reader.result)
            }
        }))
    }
   
    onFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if(e.currentTarget.files && e.currentTarget.files.length > 0) {
            const file = e.currentTarget.files && e.currentTarget.files[0];
            this.encodeBase64(file).then((result) => {
                this.setState({
                    logo: result as string 
                })
            })
        }
    }

    render() {
        return (
            <div>
                <div className="row">
                    <Header className="header" as="h1" >
                        Suppliers
                    </Header>
                </div>
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Input
                        type="text"
                        name="supplier"
                        onChange={this.onChange}
                        value={this.state.supplier}
                        fluid
                        placeholder="Supplier Name"
                        />
                        <Input
                        type="text"
                        name="address"
                        onChange={this.onChange}
                        value={this.state.address}
                        fluid
                        placeholder="Supplier Address"
                        />
                        <Input
                        name="logo"
                        type="file"
                        onChange={this.onFileChange}
                        fluid
                        />
                        <Input type="submit" value="Submit"/>
                    </Form>
                </div>
                <div className="row">
                    <Card.Group>{this.state.items}</Card.Group>
                </div>
            </div>
        );
    }
} 

export default Supplier
