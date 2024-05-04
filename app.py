import spotipy
import json
import pandas as pd
import time
from flask import Flask, request, redirect, url_for, session, render_template, jsonify

app = Flask(__name__)

app.secret_key = json.loads(open('app.json').read())['app_secret']
app.config['SESSION_COOKIE_NAME'] = 'spotify_session'


@app.route('/', methods=['GET', 'POST'])
def login():
    sp_oauth = create_spotify_oauth()
    auth_url = sp_oauth.get_authorize_url()
    print(auth_url)
    return redirect(auth_url)


@app.route('/logout')
def logout():
    for key in list(session.keys()):
        session.pop(key)
    return redirect('/')


@app.route('/authorize')
def authorize():
    sp_oauth = create_spotify_oauth()
    session.clear()
    code = request.args.get('code')
    token_info = sp_oauth.get_access_token(code)
    session['token_info'] = token_info
    return redirect("/index")


@app.route('/index', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        term = request.form.get('term')
        amount = request.form.get('amount')
        if request.form.get('get_top_tracks') == 'Top Tracks':
            return redirect(url_for('top_tracks', term=term, amount=amount))
        elif request.form.get('get_top_artists') == 'Top Artists':
            return redirect(url_for('top_artists', term=term, amount=amount))
    else:
        return render_template('/index.html')


@app.route('/top_tracks', methods=['GET', 'POST'])
def top_tracks():
    if request.method == 'POST':
        if request.form.get("back") == "Back":
            return redirect(url_for('index'))
    else:
        session['token_info'], authorized = get_token()
        session.modified = True

        # We are controlling the user if he is authorized or not.
        if not authorized:
            return redirect('/')

        # We are getting the top tracks of the user according to their input.
        time_range = request.args.get('term')
        amount_of_items = int(request.args.get('amount'))
        sp = spotipy.Spotify(auth=session['token_info']['access_token'])
        user_top_tracks  = sp.current_user_top_tracks(limit=5, time_range="long_term")

        top_tracks = {"track": [], "album": [], "artist": [], "ID": [], "popularity": [], "release_date": [],
                      "duration_ms": [], "artist_id": []}

        for i in user_top_tracks["items"]:
            top_tracks["track"].append(i['name'])
            top_tracks["album"].append(i['album']['name'])
            top_tracks["artist"].append(i['artists'][0]['name'])
            top_tracks["ID"].append(i['id'])
            top_tracks["popularity"].append(i['popularity'])
            top_tracks["release_date"].append(i['album']['release_date'])
            top_tracks["duration_ms"].append(i['duration_ms'])
            top_tracks["artist_id"].append(i['artists'][0]['id'])

            df = pd.DataFrame.from_dict(top_tracks)
            audio_analysis = sp.audio_features(df['ID'].tolist())
            df2 = pd.DataFrame.from_dict(audio_analysis)
            df_full = pd.merge(df, df2, how='inner', left_on='ID', right_on='id')
        genres = []

        for artist in df_full['artist_id']:
            genres.append(sp.artist(artist)['genres'])

        df_full['genres'] = genres

        df_full.head()
        print(df)
        print(df_full)
        print("hello")
        df_full.to_csv('output.csv', index=False)
        return render_template('/list.html', data=[])


@app.route('/top_artists', methods=['GET', 'POST'])
def top_artists():
    if request.method == 'POST':
        if request.form.get("back") == "Back":
            return redirect(url_for('index'))
    else:
        session['token_info'], authorized = get_token()
        session.modified = True

        if not authorized:
            return redirect('/')

        time_range = request.args.get('term')
        amount_of_items = int(request.args.get('amount'))
        sp = spotipy.Spotify(auth=session['token_info']['access_token'])
        list = sp.current_user_top_artists(
            limit=amount_of_items, offset=0, time_range=time_range)['items']

        result = []
        for i in range(amount_of_items):
            if i == len(list):
                break
            result.append(str(i+1) + '. ' + list[i]['name'])
        return render_template('/list.html', list=result)


def get_token():
    token_valid = False
    token_info = session.get('token_info', {})

    if not session.get('token_info', False):
        token_valid = False
        return token_info, token_valid

    now = int(time.time())
    is_token_expired = session.get(
        'token_info').get('expires_at', 0) - now < 60

    if is_token_expired:
        sp_oauth = create_spotify_oauth()
        token_info = sp_oauth.refresh_access_token(
            session.get('token_info').get('refresh_token'))

    token_valid = True
    return token_info, token_valid


def create_spotify_oauth():
    sp_oauth = spotipy.oauth2.SpotifyOAuth(
        client_id=json.loads(open('app.json').read())['client_id'],
        client_secret=json.loads(open('app.json').read())['client_secret'],
        redirect_uri=url_for('authorize', _external=True),
        scope='user-top-read',
    )
    return sp_oauth
