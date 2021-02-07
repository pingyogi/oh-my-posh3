package main

type aws struct {
	props   *properties
	env     environmentInfo
	Profile string
	Region  string
}

func (a *aws) init(props *properties, env environmentInfo) {
	a.props = props
	a.env = env
}

func (a *aws) enabled() bool {
	a.Profile = a.env.getenv("AWS_PROFILE")
	a.Region = a.env.getenv("AWS_DEFAULT_REGION")
	return a.Profile != ""
}

func (a *aws) string() string {
	segmentTemplate := a.props.getString(SegmentTemplate, "{{.Profile}}{{if .Region}}@{{.Region}}{{end}}")
	template := &textTemplate{
		Template: segmentTemplate,
		Context:  a,
	}
	return template.render()
}
