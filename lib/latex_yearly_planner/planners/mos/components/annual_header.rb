# frozen_string_literal: true

module LatexYearlyPlanner
  module Planners
    module Mos
      module Components
        class AnnualHeader < Component
          def generate(...)

            <<~LATEX
              \\marginnote{%
                #{quarters}%
              }
            LATEX
              .strip
          end

          private

          def quarters
            XTeX::VerticalStick.new(items: quarter_names, **params.object(:quarter_navigation))
          end

          def quarter_names
            params.quarters.map do |quarter|
              "#{i18n.t('calendar.one_letter.quarter')}#{quarter.number}"
            end
          end
        end
      end
    end
  end
end
