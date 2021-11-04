# frozen_string_literal: true

module Svix
    class MessageAttemptAPI
        def initialize(api_client)
            @api = MessageAttemptApi.new(api_client)
        end

        def list(app_id, msg_id, options = {})
            return @api.list_attempts_api_v1_app_app_id_msg_msg_id_attempt_get(app_id, msg_id, options)
        end

        def get(app_id, msg_id, attempt_id)
            return @api.get_attempt_api_v1_app_app_id_msg_msg_id_attempt_attempt_id_get(attempt_id, msg_id, app_id)
        end

        def resend(app_id, msg_id, endpoint_id)
            return @api.resend_webhook_api_v1_app_app_id_msg_msg_id_endpoint_endpoint_id_resend_post(endpoint_id, msg_id, app_id)
        end

        def list_attempted_messages(app_id, endpoint_id, options = {})
            return @api.list_attempted_messages_api_v1_app_app_id_endpoint_endpoint_id_msg_get(endpoint_id, app_id, options)
        end

        def list_attempted_destinations(app_id, msg_id, options = {})
            return @api.list_attempted_destinations_api_v1_app_app_id_msg_msg_id_endpoint_get(msg_id, app_id, options)
        end

        def list_attempts_for_endpoint(app_id, endpoint_id, msg_id, options = {})
            return @api.list_attempts_for_endpoint_api_v1_app_app_id_msg_msg_id_endpoint_endpoint_id_attempt_get(msg_id, app_id, endpoint_id, options)
        end
    end
end