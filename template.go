package nbcsports_replays

var indexStr = `
<!doctype html>
<html>
<head>
	<title>NBC Replays</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

	<style>
	li, p {
		font-size: 16px;
	}
	body {
		margin-bottom: 40px;
	}
	</style>

</head>
<body>
	<div class="container">
		<div class="row">
			<div class="col-md-12">
				<h1>NBC Replays</h1>
			</div>
		</div>
		<div class="row">
			<div class="col-md-8">
				<p>
				It's really difficult to get to the full event replays on the NBC
				site without seeing spoilers. Instead, all of the replays are
				listed here, with direct links to the page and no spoilers.
				</p>
				<p>
				<b>Why isn't my game here?</b> NBC doesn't post replays for all
				games, occasionally they post late, and they disappear after
				some amount of time. Sorry :(
				</p>
			</div>
		</div>
		<div class="row">
			<div class="col-md-6">
				{{ $sports := . }}
				{{ range $index, $sport := $sports }}

					{{ if ($sports.ShouldSplit $index) }}
			</div>
			<div class="col-md-6">
					{{ end }}

					<h2>{{ $sport.Name }}</h2>
					<ul>
					{{ range $sport.Events }}
						<li><a href="{{ .DestinationURL }}">{{ .Title }}</a> ({{ (.Date.Format "January 2") }})
					{{ end }}
					</ul>
					{{ if eq $sport.Name "Premier League" }}
					<p>
					Be sure to click the "Don't show scores from other games" bar
					at the top of the page, once you get there.
					{{ end }}
				{{ end }}
			</div>
		</div>
	</div>
</html>
`
