package amqptracer

import (
	opentracing "github.com/opentracing/opentracing-go"
)

func Inject(span opentracing.Span, applicationProperties map[string]interface{}) error {
	c := amqpHeadersCarrier(applicationProperties)
	return span.Tracer().Inject(span.Context(), opentracing.TextMap, c)
}

func Extract(applicationProperties map[string]interface{}) (opentracing.SpanContext, error) {
	c := amqpHeadersCarrier(applicationProperties)
	return opentracing.GlobalTracer().Extract(opentracing.TextMap, c)
}
