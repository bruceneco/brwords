package main

import (
	"fmt"
	"github.com/bruceneco/brwords"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/neilotoole/jsoncolor"
	"github.com/urfave/cli/v3"
	"os"
	"strings"
)

func formatTxtOutput(word brwords.Word, cmd *cli.Command) string {
	var sb strings.Builder
	sb.WriteString(color.BlueString("🔤 Palavra: ") + color.YellowString(word.Content))

	contentSpecified := cmd.Bool("sy") || cmd.Bool("ex") || cmd.Bool("et")
	if cmd.Bool("meaning") || !contentSpecified {
		formatMeanings(word, &sb)
	}

	if cmd.Bool("sy") {
		formatSynonyms(word, &sb)
	}

	if cmd.Bool("et") {
		formatEtymologies(word, &sb)
	}

	if cmd.Bool("ex") {
		formatExamples(word, &sb)
	}
	sb.WriteString("\n")
	return sb.String()
}

func formatExamples(word brwords.Word, sb *strings.Builder) {
	sb.WriteString(color.BlueString("\n💬 Frases de Exemplo: "))
	if len(word.Phrases) > 0 {
		for i, phrase := range word.Phrases {
			sb.WriteString(fmt.Sprintf("\n\t%d. \"%s\"", i+1, phrase.Content))
			if phrase.By != "" {
				sb.WriteString(fmt.Sprintf(" (%s)", phrase.By))
			}
		}
	} else {
		sb.WriteString(color.RedString("Nenhuma encontrada."))
	}
}

func formatEtymologies(word brwords.Word, sb *strings.Builder) {
	sb.WriteString(color.BlueString("\n🌿 Etimologias: "))
	if len(word.Etymologies) > 0 {
		for i, etymology := range word.Etymologies {
			sb.WriteString(fmt.Sprintf("\n\t%d. %s", i+1, etymology))
		}
	} else {
		sb.WriteString(color.RedString("Nenhuma encontrada."))
	}
}

func formatSynonyms(word brwords.Word, sb *strings.Builder) {
	sb.WriteString(color.BlueString("\n🔄 Sinônimos: "))
	if len(word.Synonyms) > 0 {
		sb.WriteString(fmt.Sprintf("%s", strings.Join(word.Synonyms, ", ")))
	} else {
		sb.WriteString(color.RedString("Nenhum encontrado."))
	}
}

func formatMeanings(word brwords.Word, sb *strings.Builder) {
	sb.WriteString(color.BlueString("\n📚 Significados: "))
	if len(word.Meanings) > 0 {
		for i, meaning := range word.Meanings {
			var tag string
			if meaning.Tag != "" {
				tag = fmt.Sprintf("[%s] ", meaning.Tag)
			}
			sb.WriteString(fmt.Sprintf("\n\t%d. %s%s", i+1, tag, meaning.Content))
		}
	} else {
		sb.WriteString(color.RedString("Nenhum encontrado."))
	}
}

func loadJsonEncoder() *jsoncolor.Encoder {
	var enc *jsoncolor.Encoder
	if jsoncolor.IsColorTerminal(os.Stdout) {
		out := colorable.NewColorable(os.Stdout)
		enc = jsoncolor.NewEncoder(out)
		enc.SetIndent("", "  ")
		clrs := jsoncolor.DefaultColors()
		enc.SetColors(clrs)
	} else {
		enc = jsoncolor.NewEncoder(os.Stdout)
	}
	return enc
}
