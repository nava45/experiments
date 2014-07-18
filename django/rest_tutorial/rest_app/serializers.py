from django.contrib.auth.models import User, Group
from rest_framework import serializers
from django.forms import widgets
from rest_app.models import Snippet,LANGUAGE_CHOICES, STYLE_CHOICES

class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ('url', 'username','email','groups')
        lookup_field = 'username'
        
class GroupSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Group
        fields = ('url','name')
        lookup_field = "name"
        
class SnippetSerializer(serializers.Serializer):
    pk = serializers.Field()
    title = serializers.CharField(required=False, max_length=50)
    code = serializers.CharField(widget=widgets.Textarea, max_length=100000)
    linenos = serializers.BooleanField(required=False)
    language = serializers.ChoiceField(choices=LANGUAGE_CHOICES,default="python")
    style = serializers.ChoiceField(choices=STYLE_CHOICES,
                                    default='friendly')
    
    def restore_object(self, attrs, instance=None):
        """
        Create or update a new snippet instance, given a dictionary
        of deserialized field values.

        Note that if we don't define this method, then deserializing
        data will simply return a dictionary of items.
        """
        if instance:
            # Update existing instance
            instance.title = attrs.get('title', instance.title)
            instance.code = attrs.get('code', instance.code)
            instance.linenos = attrs.get('linenos', instance.linenos)
            instance.language = attrs.get('language', instance.language)
            instance.style = attrs.get('style', instance.style)
            return instance

        # Create new instance
        return Snippet(**attrs)