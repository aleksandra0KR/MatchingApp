import numpy as np
import spotipy
from sklearn.preprocessing import StandardScaler
from spotipy.oauth2 import SpotifyClientCredentials
import pandas as pd
from sqlalchemy import create_engine
from sklearn.neighbors import KNeighborsClassifier

client_id = 'YOUR_CLIENT_ID'
client_secret = 'YOUR_SECRET'
client_credentials_manager = SpotifyClientCredentials(client_id=client_id, client_secret=client_secret)
sp = spotipy.Spotify(client_credentials_manager=client_credentials_manager)


# function of getting 50 songs from a user's playlist by username and playlist_id
def get_user_playlist(username, playlist_id):
    offset = 0
    playlist_songs = sp.user_playlist_tracks(username, playlist_id, limit=50, fields=None, offset=offset, market=None)[
        'items']
    return playlist_songs


# function of creating dataframe for new user based on her/his songs
def create_dataframe(playlist_songs):
    playlist_df = pd.DataFrame()

    playlist_df['id'] = np.array([song['track']['id'] for song in playlist_songs])
    playlist_df['user_id'] = np.array([song['added_by']['id'] for song in playlist_songs])

    audio_analysis = sp.audio_features(playlist_df['id'])

    playlist_df['danceability'] = np.array([audio_info['danceability'] for audio_info in audio_analysis])
    playlist_df['energy'] = np.array([audio_info['energy'] for audio_info in audio_analysis])
    playlist_df['key'] = np.array([audio_info['key'] for audio_info in audio_analysis])
    playlist_df['loudness'] = np.array([audio_info['loudness'] for audio_info in audio_analysis])
    playlist_df['mode'] = np.array([audio_info['mode'] for audio_info in audio_analysis])
    playlist_df['speechiness'] = np.array([audio_info['speechiness'] for audio_info in audio_analysis])
    playlist_df['acousticness'] = np.array([audio_info['acousticness'] for audio_info in audio_analysis])
    playlist_df['instrumentalness'] = np.array([audio_info['instrumentalness'] for audio_info in audio_analysis])
    playlist_df['liveness'] = np.array([audio_info['liveness'] for audio_info in audio_analysis])
    playlist_df['valence'] = np.array([audio_info['valence'] for audio_info in audio_analysis])
    playlist_df['tempo'] = np.array([audio_info['tempo'] for audio_info in audio_analysis])
    playlist_df['duration_ms'] = np.array([song['track']['duration_ms'] for song in playlist_songs])
    playlist_df['time_signature'] = np.array([audio_info['time_signature'] for audio_info in audio_analysis])

    playlist_df

    return playlist_df


def process_user_data(username, playlist_id, user_id):
    m_playlist = get_user_playlist(username, playlist_id)
    user = create_dataframe(m_playlist)
    user = user.drop(columns=['id'], axis=1)
    user = user.groupby('user_id').mean()
    user = user.reset_index()
    user = user.rename(columns={'user_id': 'user_name'})
    user1 = user.drop(columns=['user_name'], axis=1)

    # create a connection to the PostgreSQL database
    engine = create_engine('postgresql://postgres:postgres@localhost:5436/postgres')

    # write a SQL query to select data from a table
    query = 'SELECT * FROM playlists'

    # use pandas to execute the query and get a DataFrame
    df_songs = pd.read_sql_query(query, engine)

    # Load the data into a dataframe

    # Set the maximum number of rows and columns to display
    pd.set_option('display.max_rows', None)
    pd.set_option('display.max_columns', None)

    X = df_songs.drop(['user_name', 'id', 'user_id', 'playlist_key'], axis=1)
    y = df_songs['user_name']
    X = pd.concat([X, user1], ignore_index=False)

    scaler = StandardScaler()
    scaler.fit(X)
    scaled_features = scaler.transform(X)
    X = pd.DataFrame(scaled_features, columns=X.columns)
    print(X)
    user1 = X.iloc[-1:]
    X.drop(X.tail(1).index, inplace=True)

    X = X.reset_index(drop=True)

    knn = KNeighborsClassifier()
    knn.fit(X, y)

    # search for a person with a similar taste for an input user
    pred_knn = knn.predict(user1)

    user1['user_name'] = username
    user['playlist_key'] = playlist_id
    user['user_id'] = user_id
    user.to_sql('playlists', engine, if_exists='append', index=False)

    return pred_knn[0]
