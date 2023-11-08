from googletrans import Translator
from abc import ABC, abstractmethod

# Singleton Design Pattern for TranslatorFactory
class TranslatorFactory:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(TranslatorFactory, cls).__new__(cls)
            cls._instance.translators = {}
        return cls._instance

    def get_translator(self, language):
        if language not in self.translators:
            self.translators[language] = GoogleTranslator(language)
        return self.translators[language]

# Strategy Design Pattern: TranslatorStrategy interface
class TranslatorStrategy(ABC):
    @abstractmethod
    def translate(self, text):
        pass

# Adapter Design Pattern: TranslatorAdapter to standardize translation response
class TranslatorAdapter:
    def __init__(self, translator):
        self.translator = translator

    def translate(self, text):
        translation = self.translator.translate(text)
        return f"{self.translator.language}: {translation}"

# Concrete translator using the googletrans library
class GoogleTranslator(TranslatorStrategy):
    def __init__(self, language):
        self.language = language

    def translate(self, text):
        return google_translate(text, self.language)

# Replace with the googletrans library
def google_translate(text, language):
    translator = Translator()
    translation = translator.translate(text, src='auto', dest=language)
    return translation.text

if __name__ == "__main__":
    factory = TranslatorFactory()

    phrase = input("Enter a phrase or word: ")

    languages = ["ru", "en", "ja", "fr", "es"]
    language_names = ["Russian", "English", "Japanese", "French", "Spanish"]

    for lang, language in zip(languages, language_names):
        translator = factory.get_translator(lang)
        adapter = TranslatorAdapter(translator)
        translation = adapter.translate(phrase)
        print(translation)
