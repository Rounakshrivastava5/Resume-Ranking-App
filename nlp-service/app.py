from flask import Flask, request, jsonify
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
import os

app = Flask(__name__)

@app.route("/rank", methods=["POST"])
def rank_resumes():
    data = request.get_json()
    jd = data.get("job_description")
    resumes = data.get("resumes")

    if not jd or not resumes:
        return jsonify({"error": "Missing job description or resumes"}), 400

    documents = [jd] + resumes
    vectorizer = TfidfVectorizer()
    tfidf_matrix = vectorizer.fit_transform(documents)
    similarity_scores = cosine_similarity(tfidf_matrix[0:1], tfidf_matrix[1:]).flatten()

    ranked = sorted(zip(resumes, similarity_scores), key=lambda x: x[1], reverse=True)

    return jsonify([
        {"resume_text": r[0][:100], "score": round(float(r[1]), 4)} for r in ranked
    ])

if __name__ == "__main__":
    port = int(os.environ.get("PORT", 5000))
    app.run(host="0.0.0.0", port=port)
