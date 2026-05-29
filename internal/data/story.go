package data

// Chapter represents one chapter of the story.
type Chapter struct {
	Number int
	Title  string
	Body   string
	Poem   string // optional closing poem/quote
}

// Story holds the full story metadata and chapters.
type Story struct {
	Title    string
	Subtitle string
	Date     string
	Tag      string
	Chapters []Chapter
}

// LoadStory returns the full hardcoded story.
func LoadStory() *Story {
	return &Story{
		Title:    "The Tale of the Fool",
		Subtitle: "a love that was never returned, and a silence that never left.",
		Date:     "30 Dec 2025",
		Tag:      "personal story",
		Chapters: []Chapter{
			{
				Number: 1,
				Title:  "The Beginning You Don't Notice",
				Body: `It doesn't start loudly.

You don't wake up one day and decide you like someone.
It happens quietly. Slowly.

In small moments.
In conversations that didn't seem important at the time.

And before you realize it,
someone becomes part of your thoughts more than they should.

There was a moment.
A simple one.

She asked him to meet.

To anyone else, it would mean nothing.
But when you're already leaning toward someone, even slightly,
you start giving meaning to things that were never meant to carry any.

And that's how it begins.
Not with love.
But with belief.`,
				Poem: `"I was not expecting you nor was my soul searching for anyone yet,
you showed up as if the moon had whispered my need,
you aren't just a part of my life, you are the life itself."`,
			},
			{
				Number: 2,
				Title:  "The Truth You Already Know",
				Body: `You don't confess because you're ready.

You confess because you can't carry it anymore.

So he told her.

Not perfectly.
Not in the way stories make it look beautiful.

Just honestly.

And sometimes, honesty doesn't change anything.

She said she wasn't interested in this generation's love.

No confusion.
No mixed signals.

Just a clear answer.

And the worst part about a clear answer is
you can't fight it.

There's nothing to question.
Nothing to hold onto.

Just acceptance.`,
			},
			{
				Number: 3,
				Title:  "The Things That Never Happened",
				Body: `You think it's the big moments that stay.

It's not.

It's the small ones that never happened.

There was something he wanted to do.
A chocolate.

Not something important.
Not something that should matter.

But somehow, it did.

Because it wasn't about what it was.
It was about what it meant.

Later, he told her he couldn't give it.
He couldn't fulfill that small wish.

She said okay.

And sometimes, a simple "okay"
feels heavier than a goodbye.`,
				Poem: `"The writer's last wish,
As I lay down my pen, for the final time,
My heart bleeds with the words, that could never be mine.
I wrote for you, my love, my every waking thought,
But alas, my love was not enough, to make you stay.
Forever in my heart, a love that was never gained."`,
			},
			{
				Number: 4,
				Title:  "The Leaving That Looks Like Strength",
				Body: `You don't always leave because you're strong.

Sometimes you leave
because staying hurts more.

So he unfollowed her.

No message.
No drama.

Just distance.

And from the outside,
it looks like moving on.

But inside,
things don't end that easily.`,
			},
			{
				Number: 5,
				Title:  "The Way It Comes Back",
				Body: `You think you're okay.

Until something small reminds you.

A face.
A voice.
A detail you weren't expecting.

For him, it was something simple.
A girl passing by.

Nothing special.

Except one thing.
Her hair.

And just like that, everything returns.

Because memory doesn't ask you if you're ready.
It just shows up.`,
			},
			{
				Number: 6,
				Title:  "The Dream You Can't Explain",
				Body: `And then come the dreams.

The kind that don't feel like dreams.
The kind that stay even after you wake up.

He saw her.

She was crying.

Not softly.
Not in a way you can ignore.

The kind of crying that makes you feel helpless just by looking.

She held his hands tightly.
As if letting go wasn't an option.

He couldn't clearly see her face at first.
But he didn't need to.

You always know.

And she kept repeating the same thing—
That she shouldn't have said that.
Again.
And again.

And no matter how many times you hear it,
you still don't understand what it means.`,
			},
			{
				Number: 7,
				Title:  "The Version That Never Existed",
				Body: `Then comes the other kind of dream.

The dangerous one.
The one where everything is okay.

It felt like a school festival.
Noise. Laughter. Life.

And she was there.

Not distant.
Not gone.

Just there.
Like nothing ever ended.
Like everything worked out.

And that's when he saw it—
Her name.
On him.
Like something permanent.
Like something that was always meant to stay.`,
				Poem: `"They said, once he loved her so much,
he killed his own throne and bled to keep her alive."`,
			},
			{
				Number: 8,
				Title:  "The Truth You Avoid",
				Body: `You don't wake up confused.

You wake up quiet.

Because somewhere inside,
you already know the truth.

It's not about her anymore.

It's about what you thought it could be.

The version of her.
The version of you.
The version of everything.

That never became real.`,
				Poem: `"She was the moon and he was the stars,
shining brighter together yet forever apart.
He always shone through her light,
but never got to reach her,
to hold her,
to keep her to himself."`,
			},
			{
				Number: 9,
				Title:  "The Aftermath That Stays",
				Body: `It doesn't hurt the same way anymore.

But it doesn't leave either.

It just… changes.

Becomes quieter.
Heavier in a different way.

And sometimes,
it comes back in moments you didn't expect.

Now, loving someone again feels different.

Not impossible.

Just… paused.

Because once,
you loved honestly.

And that changes how you look at everything after.`,
			},
			{
				Number: 10,
				Title:  "The Words That Remain",
				Body: `In the end,
you don't go back.

You don't fix it.

You don't get closure.

You just learn to live with it.

And sometimes,
you write.

Not for them.

But because it's the only place
those feelings can exist without being questioned.`,
				Poem: `"In the end, I'm writing the verse
I once meant to share with you,
but time has changed, and now
I'm writing your presence into my poems,
through words that speak directly to my soul."`,
			},
			{
				Number: 11,
				Title:  "The First Time You Notice",
				Body: `It was an evening.

Around 7… or maybe 8.

The kind of evening that feels ordinary when it begins.

A friend had invited him for dinner.
There was a festival, people coming and going,
chairs and tables set outside,
voices blending into the night.

Nothing felt special.

Until he saw her.

For the first time.

She was just there.
Standing. Existing.

And somehow, that was enough.

He didn't think much of it.
At least, that's what he told himself.

But somewhere, quietly,
he wished to see her again.
Just once more.`,
				Poem: `"Some people don't enter your life loudly,
they arrive like a thought,
and before you notice,
they stay longer than they should."`,
			},
			{
				Number: 12,
				Title:  "The Second Look",
				Body: `He went inside with his friend.
They sat down and started eating.

Everything was normal.

But his mind was not.

And then she walked in.

Just to take something.
Nothing more.

But it was enough.
The second time.`,
				Poem: `"The first time feels like nothing,
the second time feels like something,
and before the third,
you already know it matters."`,
			},
			{
				Number: 13,
				Title:  "The Third Time",
				Body: `After everything, he left.

The houses were close to each other.
She was in one.
He was in another.

And just before leaving,
he saw her again.

The third time.

And sometimes, that is all it takes.

Not time.
Not logic.

Just a few moments that stay longer than they should.

That night, without understanding why,
he had already fallen.`,
				Poem: `"It doesn't take forever to feel something,
sometimes,
it only takes a moment that refuses to leave you."`,
			},
			{
				Number: 14,
				Title:  "The Night That Didn't End",
				Body: `He went home.

Sat down.

And then it began.

The thoughts.

Not slowly.
Not gently.

They rushed.

Her face.
Her presence.
Those few seconds repeating again and again.

No matter what he did,
his mind kept returning to her.`,
				Poem: `"You try to distract yourself,
but some thoughts don't listen,
they come back,
again and again,
until you stop resisting."`,
			},
			{
				Number: 15,
				Title:  "The Search",
				Body: `He didn't wait.

Somewhere between curiosity and something deeper,
he started searching.

Trying to find her.

A name.
A clue.
Anything.

And eventually, he found her.

Her profile.

And suddenly, something so simple
felt heavy.

He sent the request.
She accepted.
Followed back.

Just like that.
A small connection.`,
				Poem: `"It starts with something small,
a click, a follow, a moment that seems nothing,
but becomes everything later."`,
			},
			{
				Number: 16,
				Title:  "The First Words",
				Body: `There was no perfect start.

No special line.

Just a note.

Text me.

Something casual.
Something that didn't mean much.

Or at least, it wasn't supposed to.

And she replied.

She asked if it was for his lover.

And for a second,
he didn't know what to say.

Because sometimes,
a simple question
touches something deeper than expected.

He answered no.

Said he just wanted to see who would text.
Nothing serious.
Nothing deep.

Just words.

But even then,
something had already begun.`,
				Poem: `"Some stories don't start with love,
they start with a conversation,
and before you realize,
you're already part of it."`,
			},
		},
	}
}
