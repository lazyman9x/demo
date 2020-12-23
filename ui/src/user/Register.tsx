import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import React, {FormEvent} from 'react';
import Container from '../common/Container';
import DefaultPage from '../common/DefaultPage';
import {observable} from 'mobx';
import {observer} from 'mobx-react';
import Button from '@material-ui/core/Button';
import {Stores} from '../inject';
import {Theme, withStyles, WithStyles} from '@material-ui/core/styles';
import axios, {AxiosResponse} from 'axios';
import * as config from '../config';

const styles = (theme: Theme) => ({
    wrapper: {
        width: '300px'
    },
});

@observer
class Register extends React.Component<WithStyles<'wrapper'> & Stores<'currentUser'>> {
    @observable
    private username = '';
    @observable
    private email = '';
    @observable
    private password = '';
    @observable
    private repassword = '';

    public render() {
        const {username, email, password, repassword} = this;
        const {classes} = this.props;
        console.log("currentUser:", this.props);
        return (
            <DefaultPage title="Sign up" maxWidth={600}>
                <Grid item xs={8} style={{textAlign: 'center'}}>
                    <Container>
                        <form onSubmit={this.preventDefault} id="register-form">
                            <TextField
                                autoFocus
                                className={classes.wrapper}
                                label="Username"
                                margin="dense"
                                value={username}
                                onChange={(e) => (this.username = e.target.value)}
                            />
                            <TextField
                                autoFocus
                                className={classes.wrapper}
                                label="Email"
                                margin="dense"
                                value={email}
                                onChange={(e) => (this.email = e.target.value)}
                            />
                            <TextField
                                autoFocus
                                className={classes.wrapper}
                                type="password"
                                label="Password"
                                margin="dense"
                                value={password}
                                onChange={(e) => (this.password = e.target.value)}
                            />
                             <TextField
                                autoFocus
                                className={classes.wrapper}
                                type="password"
                                label="Confirm Password"
                                margin="dense"
                                value={repassword}
                                onChange={(e) => (this.repassword = e.target.value)}
                            />
                            <Button
                                variant="contained"
                                color="primary"
                                style={{marginLeft: 5, marginBottom: 5 , marginTop: 15, width: '100px'}}
                                href="register"
                                onClick={this.register}>
                                Register
                            </Button>
                        </form>
                    </Container>
                </Grid>
            </DefaultPage>
        );
    }

    private register = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        axios
            .create()
            .request({
                url: config.get('url') + 'user',
                method: 'POST',
                data: {username: this.username, email: this.email, password: this.password},
            })
            .then((resp: AxiosResponse<any>) => {
                console.log("resp: ", resp);
            })
            .catch(() => {
                console.log("Error occured");
            });
    };


    private preventDefault = (e: FormEvent<HTMLFormElement>) => e.preventDefault();
}
export default withStyles(styles)(Register);
