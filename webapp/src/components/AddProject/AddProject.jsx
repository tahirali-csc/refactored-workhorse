import React from 'react'

import TextField from '@material-ui/core/TextField'
import Grid from '@material-ui/core/Grid'
import { Typography } from '@material-ui/core'
import Button from '@material-ui/core/Button'

import ProjectService from '../../services/project-service'
import { Label } from '@material-ui/icons'


export default class AddProject extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            projectName: '',
            gitHubSSH: '',
            privateKey: ''
        }
    }

    handleSave = async () => {
        try {
            await new ProjectService().AddProject({
                name: this.state.projectName,
                privateKey: this.state.privateKey,
                cloneURL: this.state.gitHubSSH
            })
        } catch (ex) {
            console.log(ex)
        }
    }

    onChange = (e) => {
        this.setState({ [e.target.name]: e.target.value })
    }

    render() {
        return (
            <Grid container xs={12} spacing="2">
                <Grid item >
                    <Typography variant="h4" component="h4">Add Project</Typography>
                </Grid>
                <Grid item xs={12}>
                    <TextField fullWidth="true" name="projectName" required id="standard-required" label="Project Name"
                        value={this.state.projectName} onChange={this.onChange} />
                </Grid>

                <Grid item xs={12}>
                    <TextField fullWidth="true" name="gitHubSSH" required id="standard-disabled" label="GitHub SSH" value={this.state.gitHubSSH}
                        onChange={this.onChange} />
                </Grid>

                <Grid item xs={12}>
                    <TextField fullWidth="true" name="privateKey" required id="standard-disabled" label="Project Private Key" value={this.state.privateKey}
                        multiline rows={10} onChange={this.onChange} />
                </Grid>

                <Grid item xs={12}>
                    <Button variant="contained" onClick={this.handleSave}>Save</Button>
                </Grid>
            </Grid>
        )
    }

}