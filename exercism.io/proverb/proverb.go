/*
For want of a horseshoe nail, a kingdom was lost, or so the saying goes.

Given a list of inputs, generate the relevant proverb. For example, given the list `["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"]`, you will output the full text of this proverbial rhyme:

For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a nail.
*/

// Package proverb generates proverbs of variable length.
package proverb

// Function Proverb generates a proverb of variable length based on the size of the string array input.
func Proverb(input []string) []string {
	output := []string{}
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 {
			output = append(output, "And all for the want of a "+input[0]+".")
		} else {
			output = append(output, "For want of a "+input[i]+" the "+input[i+1]+" was lost.")
		}
	}
	return output
}
