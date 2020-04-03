package cli

import (
	survey "github.com/AlecAivazis/survey/v2"
)

// Doc is the contents of the doc
type Doc struct {
	Title          string
	FrontMatter    []string
	Introduction   []string
	Solutions      []string
	Considerations []string
	SuccessEval    []string
	Work           []string
	Deliberation   []string
	EndMatter      []string
}

var contents = []*survey.Question{
	{
		Name: "frontMatter",
		Prompt: &survey.MultiSelect{
			Message: "Select Front Matter items:",
			Options: []string{
				"Title",
				"Author(s)",
				"Team",
				"Reviewer(s)",
				"Created On",
				"Last Updated",
				"[Epic, Ticket, Issue] Task Link",
			},
		},
	},
	{
		Name: "introduction",
		Prompt: &survey.MultiSelect{
			Message: "Select Introduction sections:",
			Options: []string{
				"[Overview, Problem Description, Summary] Abstract",
				"[Glossary] Terminology",
				"[Context] Background",
				"[Product and Technical Requirements] Goals",
				"[Non-Goals] Out of Scope",
				"Future Goals",
				"Assumptions",
			},
		},
	},
	{
		Name: "solutions",
		Prompt: &survey.MultiSelect{
			Message: "Select Solutions sections:",
			Options: []string{
				"[Current Solution] Existing Design",
				"[Suggested Solution] Proposed Design ",
				"Test Plan",
				"Monitoring and Alerting Plan",
				"[Release Plan] Roll-out and Deployment Plan",
				"Rollback Plan",
				"[Alternate Solutions] Alternative Designs",
			},
		},
	},
	{
		Name: "considerations",
		Prompt: &survey.MultiSelect{
			Message: "Select Further Considerations sections:",
			Options: []string{
				"Impact on other teams",
				"Third-party services and platforms considerations",
				"Cost analysis",
				"Security considerations",
				"Privacy considerations",
				"Regional considerations",
				"Accessibility considerations",
				"Operational considerations",
				"Risks",
				"Support considerations",
			},
		},
	},
	{
		Name: "successEval",
		Prompt: &survey.MultiSelect{
			Message: "Select Success Evaluation sections:",
			Options: []string{
				"Impact",
				"Metrics",
			},
		},
	},
	{
		Name: "work",
		Prompt: &survey.MultiSelect{
			Message: "Select Work sections:",
			Options: []string{
				"Work Estimates and Timelines",
				"Prioritization",
				"Milestones",
				"Future work",
			},
		},
	},
	{
		Name: "deliberation",
		Prompt: &survey.MultiSelect{
			Message: "Select Deliberation sections:",
			Options: []string{
				"Discussion",
				"Open Questions",
			},
		},
	},
	{
		Name: "endMatter",
		Prompt: &survey.MultiSelect{
			Message: "Select End Matter sections:",
			Options: []string{
				"Related Work",
				"References",
				"Acknowledgments",
			},
		},
	}}

// SelectContent selects content
func SelectContent() (Doc, error) {
	var (
		selected Doc
		err      error
	)

	selected.Title, err = requestInput("Enter spec title:")
	if err != nil {
		return selected, err
	}

	err = survey.Ask(contents, &selected, survey.WithKeepFilter(true))
	if err != nil {
		return selected, err
	}

	return selected, err
}

func requestInput(request string) (string, error) {
	input := ""
	prompt := &survey.Input{
		Message: request,
	}

	err := survey.AskOne(prompt, &input)
	return input, err
}
