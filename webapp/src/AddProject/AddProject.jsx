import React from 'react'

import TextField from '@material-ui/core/TextField'
import { makeStyles } from '@material-ui/core/styles'
import Grid from '@material-ui/core/Grid'
import { Typography } from '@material-ui/core'

import Button from '@material-ui/core/Button'


// const useStyles = makeStyles((theme) => ({
//     root: {
//         '& .MuiTextField-root': {
//             margin: theme.spacing(1),
//             width: '25ch',
//         },
//     },
// }));


export default class AddProject extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            projectName : '',
            gitHubSSH : '',
            privateKey : ''
        }
    }

    handleSave = ()=> {
        console.log(this.state)
    }

    onChange = (e) => {
        this.setState({ [e.target.name] : e.target.value})
    }

    render() {
        return (
            <Grid container item xs={1} width="1000">
                <Typography variant="h4" component="h4">Add Project</Typography>
                <TextField name="projectName" required id="standard-required" label="Project Name" 
                    value={this.state.projectName} onChange={this.onChange} />

                <TextField name="gitHubSSH" required id="standard-disabled" label="GitHub SSH" value={this.state.gitHubSSH} 
                    onChange={this.onChange}/>

                <TextField name="privateKey" required id="standard-disabled" label="Project Private Key" value={this.state.privateKey}
                    multiline rows={10} onChange={this.onChange}/>

                <Button variant="contained" onClick={this.handleSave}>Save</Button>
            </Grid>
        )
    }

}