package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/pkg/errors"
)

func nameToFileName(input string) string {
	var output string
	for _, v := range input {
		if unicode.IsSpace(v) {
			output += "-"
		} else {
			output += strings.ToLower(string(v))
		}
	}
	return fmt.Sprintf("%s.jsx", output)
}

func generateFile(basePath, fileName, fileTemplate string, data interface{}) error {
	t := template.Must(template.New(fileName).Parse(fileTemplate))

	filePath := fmt.Sprintf("%s/%s", basePath, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "failed create file '%s'", filePath)
	}

	err = t.Execute(f, data)
	if err != nil {
		return errors.Wrapf(err, "failed write file '%s'", filePath)
	}

	err = f.Close()
	if err != nil {
		return errors.Wrapf(err, "failed to close file '%s'", filePath)
	}

	return nil
}

const templateStateLargeImports = ""

const templateStateLabelLarge = `
			<MapStateLabel
				x="{{ .Label.X }}"
				y="{{ .Label.Y }}"
				transform="{{ .Label.Transform }}"
				textY="{{ .Label.TextY }}"
				className={getMapLabelClassName(abbr)}
			>
				{abbr}
			</MapStateLabel>`

const templateStateSmallImports = `
	MapStateLabelBackground,
	getMapLabelBGClassName,`

const templateStateLabelSmall = `
			<>
				<MapStateLabelBackground
					x="{{ .LabelBackground.X }}"
					y="{{ .LabelBackground.Y }}"
					transform="{{ .LabelBackground.Transform }}"
					className={getMapLabelBGClassName(abbr)}
				/>
				<MapStateLabel
					x="{{ .Label.X }}"
					y="{{ .Label.Y }}"
					transform="{{ .Label.Transform }}"
					textY="{{ .Label.TextY }}"
					className={getMapLabelClassName(abbr)}
				>
					{abbr}
				</MapStateLabel>
			</>`

const templateStateComponentBase = `/** generated using horns-cli */
/** @jsx jsx */
import { jsx } from '@emotion/core'
import React from 'react'
import PropTypes from 'prop-types'
import { getColorVariants } from '../../../utils'

import MapState, {
	MapStateWrapper,
  MapStateLabel,%s
  getMapStateClassName,
  getMapLabelClassName,
} from './_map-state'

const abbr = '{{ .Abbr }}'

const {{ .ComponentName }} = ({
  fill,
  fillHover,
  fillActive,
  stroke,
  strokeHover,
  strokeActive,
  labelFill,
  labelFillHover,
  labelFillActive,
  showLabel,
}) => (
  <MapStateWrapper
    fill={fill}
    fillHover={fillHover}
    fillActive={fillActive}
    stroke={stroke}
    strokeHover={strokeHover}
    strokeActive={strokeActive}
    labelFill={labelFill}
    labelFillHover={labelFillHover}
    labelFillActive={labelFillActive}
	>
    <MapState
      d="{{ .Data }}"
      transform="{{ .Transform }}"
      className={getMapStateClassName(abbr)}
    />
    {showLabel && (%s
    )}
  </MapStateWrapper>
)

{{ .ComponentName }}.propTypes = {
  showLabel: PropTypes.bool,
  fill: PropTypes.oneOf(getColorVariants()),
  fillHover: PropTypes.oneOf(getColorVariants()),
  fillActive: PropTypes.oneOf(getColorVariants()),
  stroke: PropTypes.oneOf(getColorVariants()),
  strokeHover: PropTypes.oneOf(getColorVariants()),
  strokeActive: PropTypes.oneOf(getColorVariants()),
  labelFill: PropTypes.oneOf(getColorVariants()),
  labelFillHover: PropTypes.oneOf(getColorVariants()),
  labelFillActive: PropTypes.oneOf(getColorVariants()),
}

{{ .ComponentName }}.defaultProps = {
  showLabel: true,
  fill: 'primary',
  fillHover: 'primary-light',
  fillActive: 'primary-dark',
  stroke: 'neutral',
  strokeHover: 'neutral',
  strokeActive: 'neutral',
  labelFill: 'copy',
  labelFillHover: 'copy',
  labelFillActive: 'copy',
}

export default {{ .ComponentName }}
`

var templateStateComponentLarge = fmt.Sprintf(templateStateComponentBase, templateStateLargeImports, templateStateLabelLarge)
var templateStateComponentSmall = fmt.Sprintf(templateStateComponentBase, templateStateSmallImports, templateStateLabelSmall)

const indexTemplate = `export { default as Alabama } from './alabama'
export { default as Alaska } from './alaska'
export { default as AmericanSamoa } from './american-samoa'
export { default as Arizona } from './arizona'
export { default as Arkansas } from './arkansas'
export { default as California } from './california'
export { default as Colorado } from './colorado'
export { default as Connecticut } from './connecticut'
export { default as Delaware } from './delaware'
export { default as DistrictOfColumbia } from './district-of-columbia'
export { default as Florida } from './florida'
export { default as Georgia } from './georgia'
export { default as Guam } from './guam'
export { default as Hawaii } from './hawaii'
export { default as Idaho } from './idaho'
export { default as Illinois } from './illinois'
export { default as Indiana } from './indiana'
export { default as Iowa } from './iowa'
export { default as Kansas } from './kansas'
export { default as Kentucky } from './kentucky'
export { default as Louisiana } from './louisiana'
export { default as Maine } from './maine'
export { default as Maryland } from './maryland'
export { default as Massachusetts } from './massachusetts'
export { default as Michigan } from './michigan'
export { default as Minnesota } from './minnesota'
export { default as Mississippi } from './mississippi'
export { default as Missouri } from './missouri'
export { default as Montana } from './montana'
export { default as Nebraska } from './nebraska'
export { default as Nevada } from './nevada'
export { default as NewHampshire } from './new-hampshire'
export { default as NewJersey } from './new-jersey'
export { default as NewMexico } from './new-mexico'
export { default as NewYork } from './new-york'
export { default as NorthCarolina } from './north-carolina'
export { default as NorthDakota } from './north-dakota'
export { default as NorthernMarianaIslands } from './northern-mariana-islands'
export { default as Ohio } from './ohio'
export { default as Oklahoma } from './oklahoma'
export { default as Oregon } from './oregon'
export { default as Pennsylvania } from './pennsylvania'
export { default as PuertoRico } from './puerto-rico'
export { default as RhodeIsland } from './rhode-island'
export { default as SouthCarolina } from './south-carolina'
export { default as SouthDakota } from './south-dakota'
export { default as Tennessee } from './tennessee'
export { default as Texas } from './texas'
export { default as Utah } from './utah'
export { default as Vermont } from './vermont'
export { default as VirginIslands } from './virgin-islands'
export { default as Virginia } from './virginia'
export { default as Washington } from './washington'
export { default as WestVirginia } from './west-virginia'
export { default as Wisconsin } from './wisconsin'
export { default as Wyoming } from './wyoming'
`
